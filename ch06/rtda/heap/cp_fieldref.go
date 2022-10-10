package heap

import "Go-jvm/ch06/classfile"

/*
字段引用
*/

type FieldRef struct {
	MemberRef        // 字段符号
	field     *Field // 缓存解析指针（只需要第一次解析，这个字段相当于cache）
}

// 创建字段引用
func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

// 暴露给外部调用的解析字段符号引用
func (this *FieldRef) ResolvedField() *Field {
	if this.field == nil {
		this.resolvedFieldRef()
	}
	return this.field
}

// 解析字段符号引用
// 类D 想通过字段描述符号引用访问 类C 的某个字段，先要解析符号引用得到类C
func (this *FieldRef) resolvedFieldRef() {
	d := this.cp.class                                  // 子类
	c := this.ResolvedClass()                           // 父类/接口
	field := lookupField(c, this.name, this.descriptor) // 类c的字段

	if field == nil { // 字段查找失败
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) { // 类d 没有权限访问字段
		panic("java.lang.IllegalAccessError")
	}
	this.field = field
}

// 根据字段名和描述符查找字段
// 先查找是否是当前类c的字段；然后再去找是否是类c实现的接口中的字段；最后去找是否是类c继承的夫类的字段；如果都没有就返回nil
// 整个查找字段是一个递归的过程
func lookupField(c *Class, name string, descriptor string) *Field {
	for _, field := range c.fields { // 当前类的字段
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}

	for _, iface := range c.interfaces { // 接口的字段
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}

	if c.superClass != nil { // 父类的字段
		return lookupField(c.superClass, name, descriptor)
	}
	return nil // 字段查找失败
}

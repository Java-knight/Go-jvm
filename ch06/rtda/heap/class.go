package heap

import (
	"Go-jvm/ch06/classfile"
	"strings"
)

/*
类信息：使用结构体来表示将要放进方法区内的类
*/

type Class struct {
	accessFlags       uint16        // 类访问限定表示符
	name              string        // thisClassName
	superClassName    string        // 父类 name
	interfaceNames    []string      // 实现接口name
	constantPool      *ConstantPool // 常量池
	fields            []*Field      // 成员变量
	methods           []*Method     // 类方法
	loader            *ClassLoader  // 类加载器
	superClass        *Class        // 父类
	interfaces        []*Class      // 实现接口（接口是特殊的类）
	instanceSlotCount uint          // 接口槽位大小
	staticSlotCount   uint          // static关键字操作大小
	staticVars        Slots         // static槽位
}

// 创建类信息
func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

// check标识符
func (this *Class) IsPublic() bool {
	return 0 != this.accessFlags&ACC_ABSTRACT
}
func (this *Class) IsFinal() bool {
	return 0 != this.accessFlags&ACC_FINAL
}
func (this *Class) IsSuper() bool {
	return 0 != this.accessFlags&ACC_SUPER
}
func (this *Class) IsInterface() bool {
	return 0 != this.accessFlags&ACC_INTERFACE
}
func (this *Class) IsAbstract() bool {
	return 0 != this.accessFlags&ACC_ABSTRACT
}
func (this *Class) IsSynthetic() bool {
	return 0 != this.accessFlags&ACC_SYNTHETIC
}
func (this *Class) IsAnnotation() bool {
	return 0 != this.accessFlags&ACC_ANNOTATION
}
func (this *Class) IsEnum() bool {
	return 0 != this.accessFlags&ACC_ENUM
}

// getter method
func (this *Class) ConstantPool() *ConstantPool {
	return this.constantPool
}
func (this *Class) StaticVars() Slots {
	return this.staticVars
}

// jvm 5.4.4
// 检查 类other 是否有权限访问 类this，如果有就返回true
// 类other 访问 类this 需要满足两个条件中的一个即可：
// （1）类this是public （2）类this 和 类other 在同一个包中
func (this *Class) isAccessibleTo(other *Class) bool {
	return this.IsPublic() || this.getPackageName() == other.getPackageName()
}

// 获取当前类的包名，比如类名是java/lang/Object，包名就是java/lang；如果类定义在默认包中，它的包名是空字符串
func (this *Class) getPackageName() string {
	if i := strings.LastIndex(this.name, "/"); i >= 0 { // java/lang/Object, return java/lang
		return this.name[:i]
	}
	return ""
}

// getting main method
func (this *Class) GetMainMethod() *Method {
	return this.getStaticMethod("main", "([Ljava/lang/String;)V")
}
func (this *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range this.methods {
		if method.IsStatic() && method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return nil
}

// 创建一个对象
func (this *Class) NewObject() *Object {
	return newObject(this)
}

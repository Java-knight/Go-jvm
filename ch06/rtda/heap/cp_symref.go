package heap

/*
符号引用
*/

// symbolic reference
type SymRef struct {
	cp        *ConstantPool // 运行时常量池指针，通过符号引用访问运行时常量池，进一步访问数据
	className string        // 类的完全限定名
	class     *Class        // 缓存解析后的类结构体指针（只需要第一次解析，这个字段相当于cache）
}

// 给外部暴露的调用方法
func (this *SymRef) ResolvedClass() *Class {
	if this.class == nil {
		this.resolvedClassRef()
	}
	return this.class
}

// 解析类符号引用, jvm5.4.3.1
/*
如果类D 通过 符号引用N 引用 类C了，就要解析N，先用D的类加载器加载C，然后检查D是否有权限访问C，如果没有抛异常
*/
func (this *SymRef) resolvedClassRef() {
	d := this.cp.class
	c := d.loader.LoadClass(this.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	this.class = c
}

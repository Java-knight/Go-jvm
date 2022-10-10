package heap

import "Go-jvm/ch06/classfile"

/*
字段符号引用(字段、普通方法、接口方法引用)
问题：在Java语言中，我们并不能在同一个类中定义名字相同，但类型不同的两个字段，那么字段符号的引用为什么还要存放字段描述符呢？
因为这只是Java语言的限制，而不是Java虚拟机的要求，在Java虚拟机一个类是完全可以有多个同名字段的，只要它们的类型互不相同就行
*/

// 运行时常量池的引用信息
type MemberRef struct {
	SymRef     // 字段描述符
	name       string
	descriptor string
}

// 从 常量池 拷贝出需要的信息到 运行时常量池（字段引用、普通方法引用、接口方法引用）
func (this *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	this.className = refInfo.ClassName()
	this.name, this.descriptor = refInfo.NameAndDescriptor()
}

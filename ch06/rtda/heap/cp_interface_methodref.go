package heap

import "Go-jvm/ch06/classfile"

/*
接口方法符号引用
*/

type InterfaceMethodRef struct {
	MemberRef         // 符号引用
	method    *Method // 缓存解析指针（只需要第一次解析，这个字段相当于cache）
}

// 创建接口方法引用
func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

// 将接口方法转换为普通方法
func (this *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if this.method == nil {
		this.resolvedInterfaceMethod()
	}
	return this.method
}

// jvms8 5.4.3.4
func (this *InterfaceMethodRef) resolvedInterfaceMethod() {
	//class := this.ResloveClass()
	//todo
}

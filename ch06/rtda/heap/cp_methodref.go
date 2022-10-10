package heap

import "Go-jvm/ch06/classfile"

/*
普通方法符号引用
*/

type MethodRef struct {
	MemberRef         // 字段符号
	method    *Method // 缓存解析指针（只需要第一次解析，这个字段相当于cache）
}

// 创建普通方法引用
func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

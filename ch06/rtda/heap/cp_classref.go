package heap

import "Go-jvm/ch06/classfile"

/*
类符号引用
*/

type ClassRef struct {
	SymRef
}

// 创建类符号引用
func newClassRef(cp *ConstantPool, classInfo *classfile.ConstantClassInfo) *ClassRef {
	ref := &ClassRef{}
	ref.cp = cp
	ref.className = classInfo.Name()
	return ref
}

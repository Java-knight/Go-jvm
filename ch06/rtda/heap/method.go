package heap

import "Go-jvm/ch06/classfile"

/*
方法
*/

type Method struct {
	ClassMember
	maxStack  uint
	maxLocals uint
	code      []byte // 字节码
}

// 创建该类所有的方法(将classfile中的信息加载到方法区中)
func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttributes(cfMethod)
	}
	return methods
}

// 拷贝属性信息
func (this *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		this.maxStack = codeAttr.MaxStack()
		this.maxLocals = codeAttr.MaxLocals()
		this.code = codeAttr.Code()
	}
}

// check synchronized/bridge/varargs/native/abstract/strict
func (this *Method) IsSynchronized() bool {
	return 0 != this.accessFlags&ACC_SYNCHRONIZED
}
func (this *Method) IsBridge() bool {
	return 0 != this.accessFlags&ACC_BRIDGE
}
func (this *Method) IsVarargs() bool {
	return 0 != this.accessFlags&ACC_VARARGS
}
func (this *Method) IsNative() bool {
	return 0 != this.accessFlags&ACC_NATIVE
}
func (this *Method) IsAbstract() bool {
	return 0 != this.accessFlags&ACC_ABSTRACT
}
func (this *Method) IsStrict() bool {
	return 0 != this.accessFlags&ACC_STRICT
}

// getters
func (this *Method) MaxStack() uint {
	return this.maxStack
}
func (this *Method) MaxLocals() uint {
	return this.maxLocals
}
func (this *Method) Code() []byte {
	return this.code
}

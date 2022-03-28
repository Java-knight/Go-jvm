package classfile

/*
CONSTANT_MethodType_info、CONSTANT_MethodHandle_info、CONSTANT_InvokeDynamic_info
它们都是Java7才添加到class文件中的, 目的是支持新增的invokedynamic 指令
*/

// 方法处理 常量信息
type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

// 方法类型 常量信息
type ConstantMethodTypeInfo struct {
	descriptor uint16
}

// 动态调用(动态代理) 常量信息
type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (this *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	this.referenceKind = reader.readUint8()
	this.referenceIndex = reader.readUint16()
}

func (this *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	this.descriptor = reader.readUint16()
}

func (this *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	this.bootstrapMethodAttrIndex = reader.readUint16()
	this.nameAndTypeIndex = reader.readUint16()
}

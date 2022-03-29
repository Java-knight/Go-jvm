package classfile

// CONSTANT_MethodType_info: https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4.9
// CONSTANT_MethodHandle_info: https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4.8
// CONSTANT_InvokeDynamic_info: https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4.10

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

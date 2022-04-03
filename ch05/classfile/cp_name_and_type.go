package classfile

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4.6

// 字段或方法 的名称和描述符
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (this *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	this.nameIndex = reader.readUint16()
	this.descriptorIndex = reader.readUint16()
}

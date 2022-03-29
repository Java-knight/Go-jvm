package classfile

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.7.2

// 定长属性: 表示常量表达式的值
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (this *ConstantValueAttribute) readInfo(reader *ClassReader) {
	this.constantValueIndex = reader.readUint16()
}

func (this *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return this.constantValueIndex
}

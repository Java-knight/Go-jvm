package classfile

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4.3

// 字符串 常量信息
type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

// 读取常量池的索引
func (this *ConstantStringInfo) readInfo(reader *ClassReader) {
	this.stringIndex = reader.readUint16()
}

// toString()打印（按照索引查找字符串）
func (this *ConstantStringInfo) String() string {
	return this.cp.getUtf8(this.stringIndex)
}

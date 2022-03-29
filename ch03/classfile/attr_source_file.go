package classfile

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.7.10

// SourceFile属性: 可选定长属性, 用于指处源文件名
type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (this *SourceFileAttribute) readInfo(reader *ClassReader) {
	this.sourceFileIndex = reader.readUint16()
}

// toString()打印
func (this *SourceFileAttribute) FileName() string {
	return this.cp.getUtf8(this.sourceFileIndex)
}

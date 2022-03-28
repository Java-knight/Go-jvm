package classfile

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

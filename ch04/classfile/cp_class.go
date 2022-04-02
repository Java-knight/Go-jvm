package classfile

// class类 常量信息
type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (this *ConstantClassInfo) readInfo(reader *ClassReader) {
	this.nameIndex = reader.readUint16()
}

// toString()打印（通过索引去找）
func (this *ConstantClassInfo) Name() string {
	return this.cp.getUtf8(this.nameIndex)
}

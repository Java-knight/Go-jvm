package classfile

// 未解析属性（没有进行预定义）
type UnparsedAttribute struct {
	name   string // 属性名称
	length uint32 // 属性长度
	info   []byte // 属性信息
}

func (this *UnparsedAttribute) readInfo(reader *ClassReader) {
	this.info = reader.readByte(this.length)
}

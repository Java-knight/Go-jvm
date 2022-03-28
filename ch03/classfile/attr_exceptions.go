package classfile

// Exceptions属性: 变长属性, 记录方法抛出的异常表
type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (this *ExceptionsAttribute) readInfo(reader *ClassReader) {
	this.exceptionIndexTable = reader.readUint16s()
}

// toString()打印
func (this *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return this.exceptionIndexTable
}

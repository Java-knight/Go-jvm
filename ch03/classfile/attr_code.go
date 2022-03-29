package classfile

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.7.3

// Code属性: 变长属性, 存放字节码等方法相关信息
type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16                 // 操作数栈最大深度
	maxLocals      uint16                 // 局部变量表大小（slot槽大小）
	code           []byte                 // 字节码
	exceptionTable []*ExceptionTableEntry // 异常表
	attributes     []AttributeInfo        // 属性
}

// 方法异常表
type ExceptionTableEntry struct {
	startPc   uint16
	entPc     uint16
	handlerPc uint16
	catchType uint16
}

func (this *CodeAttribute) readInfo(reader *ClassReader) {
	this.maxStack = reader.readUint16()
	this.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	this.code = reader.readByte(codeLength)
	this.exceptionTable = readExceptionTable(reader)
	this.attributes = readAttributes(reader, this.cp)
}

// 获取异常表
func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			entPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}

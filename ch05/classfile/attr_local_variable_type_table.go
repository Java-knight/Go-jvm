package classfile

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.7.14

// 局部变量表类型属性: 可变长属性, 因为一个虚拟机栈有多个方法
type LocalVariableTypeTableAttribute struct {
	localVariableTypeTable []*LocalVariableTypeTableEntry
}

// 局部变量表实体
type LocalVariableTypeTableEntry struct {
	startPc        uint16
	length         uint16
	nameIndex      uint16
	signatureIndex uint16 // 签名索引, 对局部变量表的类型践行编码
	index          uint16
}

func (this *LocalVariableTypeTableAttribute) readInfo(reader *ClassReader) {
	localVariableTypeTableLength := reader.readUint16()
	this.localVariableTypeTable = make([]*LocalVariableTypeTableEntry, localVariableTypeTableLength)
	for i := range this.localVariableTypeTable {
		this.localVariableTypeTable[i] = &LocalVariableTypeTableEntry{
			startPc:        reader.readUint16(),
			length:         reader.readUint16(),
			nameIndex:      reader.readUint16(),
			signatureIndex: reader.readUint16(),
			index:          reader.readUint16(),
		}
	}

}

package classfile

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.7.13

// 局部变量表属性: 可变长度属性, 一个虚拟机栈中有多个方法/栈桢, 每个方法一个局部变量表
type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableTableEntry
}

// 局部变量表实体, 表中的局部变量以任意顺序出现(slot槽可以复用, 普通方法第一行一定是this, static方法没有this)
type LocalVariableTableEntry struct {
	// 理解: 一个线程的多个栈帧局部变量表是一个数组, 这个数组是记录类对应局部变量表的起始位置和偏移量
	startPc         uint16 // 位于code array[startPc, startPc+length]之间,
	length          uint16
	nameIndex       uint16 //
	descriptorIndex uint16 // 描述符索引
	index           uint16 // 栈帧位置, long和double占2位
}

func (this *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	localVariableTableLength := reader.readUint16()
	this.localVariableTable = make([]*LocalVariableTableEntry, localVariableTableLength)
	for i := range this.localVariableTable {
		this.localVariableTable[i] = &LocalVariableTableEntry{
			startPc:         reader.readUint16(),
			length:          reader.readUint16(),
			nameIndex:       reader.readUint16(),
			descriptorIndex: reader.readUint16(),
			index:           reader.readUint16(),
		}
	}

}

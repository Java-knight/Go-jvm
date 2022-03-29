package classfile

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.7.12

// LineNumberTable属性: 存放方法的行号信息
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

// LineNumberTable属性实体（单个方法）
type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (this *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	this.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range this.lineNumberTable {
		this.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}

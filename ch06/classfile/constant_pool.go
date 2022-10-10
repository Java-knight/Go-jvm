package classfile

// 常量池
type ConstantPool []ConstantInfo

// 读取常量池信息, 将[]byte转化为ConstantPool
func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)

	for i := 1; i < cpCount; i++ { // 注意索引从1开始
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++ // long和double都是占 constant_pool两个位置
		}
	}
	return cp
}

// 	通过index 在常量池中获取对应的 ConstantInfo
func (this ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := this[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

// 从常量池查找字段或方法的名字和描述符
func (this ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := this.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := this.getUtf8(ntInfo.nameIndex)
	_type := this.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

// 从常量池查找类名
func (this ConstantPool) getClassName(index uint16) string {
	classInfo := this.getConstantInfo(index).(*ConstantClassInfo)
	return this.getUtf8(classInfo.nameIndex)
}

// 从常量池查找 UTF-8 字符串
func (this ConstantPool) getUtf8(index uint16) string {
	utf8Info := this.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}

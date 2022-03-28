package classfile

// 常量池的引用信息（字段引用、普通方法引用、接口方法引用三个的结构体相同）
type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16 // 类索引
	nameAndTypeIndex uint16 // 名字和类型（字段或方法的名称和描述符）索引
}

// 字段引用（"继承"了ConstantMemberrefInfo）
type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

// 普通方法引用（"继承"了ConstantMemberrefInfo）
type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

// 接口方法引用（"继承"了ConstantMemberrefInfo）
type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}

func (this *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	this.classIndex = reader.readUint16()
	this.nameAndTypeIndex = reader.readUint16()
}

// 类名打印(根据索引去找)
func (this *ConstantMemberrefInfo) ClassName() string {
	return this.cp.getClassName(this.classIndex)
}

// 字段或方法的 名称和描述符 打印(根据索引去找)
func (this *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return this.cp.getNameAndType(this.nameAndTypeIndex)
}

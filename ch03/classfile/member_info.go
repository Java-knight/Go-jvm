package classfile

// 字段和方法表信息
type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16          // 描述符信息
	attributes      []AttributeInfo // 成员变量
}

// 读取字段表或方法表（字段表就是槽位slot 中的局部变量）
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

// 读取成员信息（单个）
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}

// getting access_flags
func (this *MemberInfo) AccessFlags() uint16 {
	return this.accessFlags
}

// getting name
func (this *MemberInfo) Name() string {
	return this.cp.getUtf8(this.nameIndex)
}

// getting 描述符
func (this *MemberInfo) Descriptor() string {
	return this.cp.getUtf8(this.descriptorIndex)
}

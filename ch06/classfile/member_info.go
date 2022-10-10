package classfile

// 字段：https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.5
// 方法：https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.6

// 字段和方法表信息
type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16          // 描述符信息
	attributes      []AttributeInfo // 属性信息
}

// 获取成员信息
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

// 获取 字节码属性
func (this *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range this.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

// 获取 常量池中的定长属性
func (this *MemberInfo) ConstantValueAttribute() *ConstantValueAttribute {
	for _, attrInfo := range this.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}

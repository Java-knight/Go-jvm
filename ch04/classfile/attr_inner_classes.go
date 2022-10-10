package classfile

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.7.6

// 内部类属性: 可变长度属性, 如果类或接口的常量池包含至少一个CONSTANT_Class_info,
// 它表示不是包成员的类或接口, 则表示表中 必须恰好有一个InnerClasses属性
type InnerClassedAttribute struct {
	classes []*InnerClassInfo
}

// 内部类信息
type InnerClassInfo struct {
	innerClassInfoIndex   uint16 // 内部类索引
	outerClassInfoIndex   uint16 // 外部类索引
	innerNameIndex        uint16 // 内部名称索引
	innerClassAccessFlags uint16 // 用于表示对类或接口的访问权限和属性(查看文档)
}

func (this *InnerClassedAttribute) readInfo(reader *ClassReader) {
	numberOfClasses := reader.readUint16()
	this.classes = make([]*InnerClassInfo, numberOfClasses)
	for i := range this.classes {
		this.classes[i] = &InnerClassInfo{
			innerClassInfoIndex:   reader.readUint16(),
			outerClassInfoIndex:   reader.readUint16(),
			innerNameIndex:        reader.readUint16(),
			innerClassAccessFlags: reader.readUint16(),
		}
	}
}

package classfile

import "fmt"

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.1

// Class文件格式（常量池、方法信息、接口信息、属性信息）
type ClassFile struct {
	magic        uint32       // 魔术
	minorVersion uint16       // JDK次版本号
	majorVersion uint16       // JDK主版本号
	constantPool ConstantPool // 常量池
	accessFlags  uint16       // 类访问限定标识符
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo   // 成员变量
	methods      []*MemberInfo   // 成员方法
	attributes   []AttributeInfo // 属性信息（局部变量表、异常信息、符号信息...）
}

// 解析字节码（将[]byte 转换成 ClassFile结构体）
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

// 读取class字节流(使用ClassReader按照数据类型进行分类)
func (this *ClassFile) read(reader *ClassReader) {
	this.readAndCheckMagic(reader)
	this.readAndCheckVersion(reader)
	this.constantPool = readConstantPool(reader)
	this.accessFlags = reader.readUint16()
	this.thisClass = reader.readUint16()
	this.superClass = reader.readUint16()
	this.interfaces = reader.readUint16s()
	this.fields = readMembers(reader, this.constantPool)
	this.methods = readMembers(reader, this.constantPool)
	this.attributes = readAttributes(reader, this.constantPool)
}

// 读取和校验 magic(魔术), 校验是否是class文件
func (this *ClassFile) readAndCheckMagic(reader *ClassReader) {
	this.magic = reader.readUint32()
	if this.magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}

}

// 读取和校验 version（校验jdk版本, jdk1.6 对应 魔术版本50, 1.7对应51...）
func (this *ClassFile) readAndCheckVersion(reader *ClassReader) {
	this.minorVersion = reader.readUint16()
	this.majorVersion = reader.readUint16()
	switch this.majorVersion {
	case 45:
		return
	case 46, 48, 49, 50, 51, 52:
		if this.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

// getting minor_version
func (this *ClassFile) MinorVersion() uint16 {
	return this.minorVersion
}

// getting major_version
func (this *ClassFile) MajorVersion() uint16 {
	return this.majorVersion
}

// getting constant_pool
func (this *ClassFile) ConstantPool() ConstantPool {
	return this.constantPool
}

// getting access_flags
func (this *ClassFile) AccessFlags() uint16 {
	return this.accessFlags
}

// getting fields
func (this *ClassFile) Fields() []*MemberInfo {
	return this.fields
}

// getting method
func (this *ClassFile) Methods() []*MemberInfo {
	return this.methods
}

// getting className
func (this *ClassFile) ClassName() string {
	return this.constantPool.getClassName(this.thisClass)
}

// getting superName
func (this *ClassFile) SuperClassName() string {
	if this.superClass > 0 {
		return this.constantPool.getClassName(this.superClass)
	}
	return "" // 只有 java.lang.Object没有超类
}

// getting interface names
func (this *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(this.interfaces))

	for i, cpIndex := range this.interfaces {
		interfaceNames[i] = this.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

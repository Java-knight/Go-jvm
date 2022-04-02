package classfile

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.7.7

// 封闭方法属性: 固定长度，当前仅当 类表示本地类或者匿名类, 类必须要有封闭方法属性, 类最大只能有一个 封闭方法属性
type EnclosingMethodAttribute struct {
	cp          ConstantPool
	classIndex  uint16
	methodIndex uint16
}

func (this *EnclosingMethodAttribute) readInfo(reader *ClassReader) {
	this.classIndex = reader.readUint16()
	this.methodIndex = reader.readUint16()
}

// toString()打印类名
func (this *EnclosingMethodAttribute) ClassName() string {
	return this.cp.getClassName(this.classIndex)
}

// 打印方法名字和描述符
func (this *EnclosingMethodAttribute) MethodNameAndDescriptor() (string, string) {
	if this.methodIndex > 0 {
		return this.cp.getNameAndType(this.methodIndex)
	} else {
		return "", ""
	}
}

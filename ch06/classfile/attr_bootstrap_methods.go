package classfile

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.7.23

// 启动方法属性: invokedynamic指令的引导方法说明
type BootstrapMethodsAttribute struct {
	bootstrapMethods []*BootstrapMethod
}

func (this *BootstrapMethodsAttribute) readInfo(reader *ClassReader) {
	numBootstrapMethods := reader.readUint16()
	this.bootstrapMethods = make([]*BootstrapMethod, numBootstrapMethods)
	for i := range this.bootstrapMethods {
		this.bootstrapMethods[i] = &BootstrapMethod{
			bootstrapMethodRef: reader.readUint16(),
			bootstrapArguments: reader.readUint16s(),
		}
	}
}

// 启动方法
type BootstrapMethod struct {
	bootstrapMethodRef uint16   // 启动方法引用
	bootstrapArguments []uint16 // 启动方法参数
}

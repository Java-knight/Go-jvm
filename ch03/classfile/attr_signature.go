package classfile

// https://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.7.9

// 签名属性: 记录类、接口、构造函数、方法或字段的签名
// 签名: 就是对Java编写代码的一种声明式编码, 比如B C D F I J S Z就是对一些基本类型的签名, 还有一些参数类型,包等等
// 详细的签名类型请查看官方文档
type SignatureAttribute struct {
	cp             ConstantPool
	signatureIndex uint16
}

func (this *SignatureAttribute) readInfo(reader *ClassReader) {
	this.signatureIndex = reader.readUint16()
}

// toString()打印
func (this *SignatureAttribute) Signature() string {
	return this.cp.getUtf8(this.signatureIndex)
}

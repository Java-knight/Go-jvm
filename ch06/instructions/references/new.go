package references

import (
	"Go-jvm/ch06/instructions/base"
	"Go-jvm/ch06/rtda"
	"Go-jvm/ch06/rtda/heap"
)

/*
new指令：Object obj = new Object(); Java中创建对象时需要使用到的指令
*/

// Create new object
type NEW struct {
	base.Index16Instruction
}

// new指令的操作数是一个 uint16 索引，来自字节码。
// 通过这个索引，可以从当前类的运行时常量池中找到一个类符号引用。
// 解析这个符号引用，拿到类数据，然后创建对象，并把对象引用推入栈顶，new指令工作完成
func (this *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(this.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()

	// Java中接口和抽象类不能初始化
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}

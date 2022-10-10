package references

import (
	"Go-jvm/ch06/instructions/base"
	"Go-jvm/ch06/rtda"
	"Go-jvm/ch06/rtda/heap"
)

/*
instanceof指令：判断对象是否是某个类的实例（或者对象的类是否实现类某个结构），并把结果推入操作数栈。需要两个操作数。
第一个操作数是 uint16 索引，从方法的字节码中获取，通过这个索引可以从当前类的运行时常量池中找到一个类符号引用。
第二个操作数是对象引用，从操作数栈中弹出。
*/

// Determine if object is of given type
type INSTANCE_OF struct {
	base.Index16Instruction
}

// 先弹出对象引用，如果是 null，则把 0 推入操作数栈。用Java代码就是就是，如果引用 obj 是 null的话，不管ClassXXX是哪种类型，下面这条if判断都是false
// if (obj instanceof ClassXXX) {...}
func (this *INSTANCE_OF) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		stack.PushInt(0)
		return
	}
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(this.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}

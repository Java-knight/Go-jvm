package references

import (
	"Go-jvm/ch06/instructions/base"
	"Go-jvm/ch06/rtda"
	"Go-jvm/ch06/rtda/heap"
)

/*
checkcast指令：和instanceof指令很像
区别在于：instanceof指令会改变操作数栈（弹出对象引用，推入判断结果）；
        checkcast则不改变操作数栈（如果判断失败，直接抛出 ClassCastException异常）
*/

// Check whether object is of given type
type CHACK_CAST struct {
	base.Index16Instruction
}

// 先从操作数栈中弹出对象引用，再堆回去，不会改变操作数栈堆状态。如果引用是 null，则指令执行结束。
// null 引用可以转换成任何类型，否则解析类符号引用，判断对象是否是类实例。如果是堆话，指令执行结束，否则抛出 ClassCastExeception
/*
if (xxx instanceof ClassYYY) {
    yyy = (ClassXXX) xxx
    // use yyy
}
*/
func (this *CHACK_CAST) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref) // 不会概念操作数栈中的内容
	if ref == nil {
		return
	}
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(this.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}

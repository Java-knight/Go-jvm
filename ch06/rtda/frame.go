package rtda

import "Go-jvm/ch06/rtda/heap"

// 栈桢的实现
// TODO 栈桢中包括局部变量表、操作数栈、方法返回地址、动态链接、一些附加信息（可有可无）
type Frame struct {
	next         *Frame        // 当前栈帧
	localVars    LocalVars     // 局部变量表local variable table
	operandStack *OperandStack // 操作数栈
	thread       *Thread       // 当前线程
	nextPC       int           // 当前程序计数器
	method       *heap.Method  // 当前方法
}

// 创建一个栈桢
// maxLocals: 局部变量表slots大小
// maxStack: 操作数栈容量大小
func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
	}
}

// getting 局部变量表
func (this *Frame) LocalVars() LocalVars {
	return this.localVars
}

// getting 操作数栈
func (this *Frame) OperandStack() *OperandStack {
	return this.operandStack
}

// getting 当前线程
func (this *Frame) Thread() *Thread {
	return this.thread
}

// getting 程序计数器
func (this *Frame) NextPC() int {
	return this.nextPC
}

// getting 方法
func (this *Frame) Method() *heap.Method {
	return this.method
}

// setting 程序计数器
func (this *Frame) SetNextPC(nextPC int) {
	this.nextPC = nextPC
}

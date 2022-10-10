package rtda

// 栈桢的实现
// TODO 栈桢中包括局部变量表、操作数栈、方法返回地址、动态链接、一些附加信息（可有可无）
type Frame struct {
	next         *Frame        // 当前栈帧
	localVars    LocalVars     // 局部变量表local variable table
	operandStack *OperandStack // 操作数栈
	thread       *Thread       // 当前线程
	nextPC       int           // 当前程序计数器
}

// 创建一个栈桢
// maxLocals: 局部变量表slots大小
// maxStack: 操作数栈容量大小
func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
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

// setting 程序计数器
func (this *Frame) SetNextPC(nextPC int) {
	this.nextPC = nextPC
}

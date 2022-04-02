package rtda

// 栈桢的实现
// TODO 栈桢中包括局部变量表、操作数栈、方法返回地址、动态链接、一些附加信息（可有可无）
type Frame struct {
	next         *Frame
	localVars    LocalVars     // 局部变量表local variable table
	operandStack *OperandStack // 操作数栈
}

// 创建一个栈桢
// maxLocals: 局部变量表slots大小
// maxStack: 操作数栈容量大小
func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
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

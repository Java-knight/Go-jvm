package rtda

// 线程
type Thread struct {
	pc    int    // PC程序计数器
	stack *Stack // 虚拟机栈/本地方法栈
}

// 创建一个线程
// TODO 这里的栈桢需要通过cmd读取到-Xss进行指定, 默认是1024
func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

// getting PC程序计数器
func (this *Thread) PC() int {
	return this.pc
}

// setting PC程序计数器
func (this *Thread) SetPC(pc int) {
	this.pc = pc
}

// push栈桢
func (this *Thread) PushFrame(frame *Frame) {
	this.stack.push(frame)
}

// pop 出一个栈桢
func (this *Thread) PopFrame() *Frame {
	return this.stack.pop()
}

// 获取栈顶栈桢
func (this *Thread) CurrentFrame() *Frame {
	return this.stack.peek()
}

// 创建一个栈帧(暴露给外部调用的)
func (this *Thread) NewFrame(maxLocals, maxStack uint) *Frame {
	return newFrame(this, maxLocals, maxStack)
}

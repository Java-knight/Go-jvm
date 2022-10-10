package rtda

// java虚拟机栈/本地方法栈（逻辑上是两个，实现上是一个，这里采用链表实现）
type Stack struct {
	maxSize uint   // 栈的最大容量
	size    uint   // 当前栈大小
	_top    *Frame // 栈顶的栈桢
}

// 创建一个栈
func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

// frame3 -> frame2 -> frame1 -> nil
//   |
//  _top
// ***********push(frame4)*************
// frame4 -> frame3 -> frame2 -> frame1 -> nil
//   |
//  _top
// 入栈操作（链表的下一个是栈顶元素）
func (this *Stack) push(frame *Frame) {
	if this.size >= this.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if this._top != nil {
		frame.next = this._top
	}
	this._top = frame
	this.size++
}

// frame4 -> frame3 -> frame2 -> frame1 -> nil
//    |
//   _top
//********************pop()***************
// frame3 -> frame2 -> frame1 -> nil     frame4
//    |                                    |
//   _top                                 top
// 出栈操作
func (this *Stack) pop() *Frame {
	if this._top == nil {
		panic("jvm stack is empty!")
	}
	top := this._top
	this._top = top.next
	top.next = nil
	this.size--
	return top
}

// 返回栈顶元素（不取出来）
func (this *Stack) peek() *Frame {
	if this._top == nil {
		panic("jvm stack is empty!")
	}
	return this._top
}

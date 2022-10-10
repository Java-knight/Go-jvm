package stack

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
定义pop和pop2指令
*/

type POP struct {
	base.NoOperandsInstruction
}

type POP2 struct {
	base.NoOperandsInstruction
}

// 出栈一个操作数栈的栈顶元素（只能用于弹出int\byte\short\chat、float等占一个操作数位置的变量）
func (this *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// 出栈一个操作数栈元素（可以弹出long、double占用两个操作数位置的变量）
func (this *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}

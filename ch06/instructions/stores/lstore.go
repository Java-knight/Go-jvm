package stores

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
lstore指令：long类型数据从操作数栈中弹出数据，存放中局部变量表中
Store long info local variable
*/

type LSTORE struct {
	base.Index8Instruction
}

type LSTORE_0 struct {
	base.NoOperandsInstruction
}

type LSTORE_1 struct {
	base.NoOperandsInstruction
}

type LSTORE_2 struct {
	base.NoOperandsInstruction
}

type LSTORE_3 struct {
	base.NoOperandsInstruction
}

func (this *LSTORE) Execute(frame *rtda.Frame) {
	_lstore(frame, uint(this.Index))
}

// LSTORE_0/LSTORE_1/LSTORE_2/LSTORE_3 的索引隐含中操作码中
func (this *LSTORE_0) Execute(frame *rtda.Frame) {
	_lstore(frame, 0)
}

func (this *LSTORE_1) Execute(frame *rtda.Frame) {
	_lstore(frame, 1)
}

func (this *LSTORE_2) Execute(frame *rtda.Frame) {
	_lstore(frame, 2)
}

func (this *LSTORE_3) Execute(frame *rtda.Frame) {
	_lstore(frame, 3)
}

// 提供一个公共函数(将栈桢的操作数栈顶的元素pop出，存储到局部变量表)
func _lstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

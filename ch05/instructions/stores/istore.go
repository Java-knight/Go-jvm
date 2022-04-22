package stores

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
istore指令：int类型数据从操作数栈中弹出数据，存放中局部变量表中
Store int info local variable
*/

type ISTORE struct {
	base.Index8Instruction
}

type ISTORE_0 struct {
	base.NoOperandsInstruction
}

type ISTORE_1 struct {
	base.NoOperandsInstruction
}

type ISTORE_2 struct {
	base.NoOperandsInstruction
}

type ISTORE_3 struct {
	base.NoOperandsInstruction
}

func (this *ISTORE) Execute(frame *rtda.Frame) {
	_istore(frame, this.Index)
}

func (this *ISTORE_0) Execute(frame *rtda.Frame) {
	_istore(frame, 0)
}

func (this *ISTORE_1) Execute(frame *rtda.Frame) {
	_istore(frame, 1)
}

func (this *ISTORE_2) Execute(frame *rtda.Frame) {
	_istore(frame, 2)
}

func (this *ISTORE_3) Execute(frame *rtda.Frame) {
	_istore(frame, 3)
}

func _istore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}

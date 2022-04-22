package stores

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
fstore指令：浮点float类型数据从操作数栈中弹出数据，存放中局部变量表中
Store float info local variable
*/

type FSTORE struct {
	base.Index8Instruction
}

type FSTORE_0 struct {
	base.NoOperandsInstruction
}

type FSTORE_1 struct {
	base.NoOperandsInstruction
}

type FSTORE_2 struct {
	base.NoOperandsInstruction
}

type FSTORE_3 struct {
	base.NoOperandsInstruction
}

func (this *FSTORE) Execute(frame *rtda.Frame) {
	_fstore(frame, this.Index)
}

func (this *FSTORE_0) Execute(frame *rtda.Frame) {
	_fstore(frame, 0)
}

func (this *FSTORE_1) Execute(frame *rtda.Frame) {
	_fstore(frame, 1)
}

func (this *FSTORE_2) Execute(frame *rtda.Frame) {
	_fstore(frame, 2)
}

func (this *FSTORE_3) Execute(frame *rtda.Frame) {
	_fstore(frame, 3)
}

func _fstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}

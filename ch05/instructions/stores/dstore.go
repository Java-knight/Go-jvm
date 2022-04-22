package stores

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
dstore指令：浮点double类型数据从操作数栈中弹出数据，存放中局部变量表中
Store double info local variable
*/

type DSTORE struct {
	base.Index8Instruction
}

type DSTORE_0 struct {
	base.NoOperandsInstruction
}

type DSTORE_1 struct {
	base.NoOperandsInstruction
}

type DSTORE_2 struct {
	base.NoOperandsInstruction
}

type DSTORE_3 struct {
	base.NoOperandsInstruction
}

func (this *DSTORE) Execute(frame *rtda.Frame) {
	_dstore(frame, this.Index)
}

func (this *DSTORE_0) Execute(frame *rtda.Frame) {
	_dstore(frame, 0)
}

func (this *DSTORE_1) Execute(frame *rtda.Frame) {
	_dstore(frame, 1)
}

func (this *DSTORE_2) Execute(frame *rtda.Frame) {
	_dstore(frame, 2)
}

func (this *DSTORE_3) Execute(frame *rtda.Frame) {
	_dstore(frame, 3)
}

func _dstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}

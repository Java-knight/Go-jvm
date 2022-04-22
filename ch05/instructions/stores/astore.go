package stores

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
astore指令：引用reference类型数据从操作数栈中弹出数据，存放中局部变量表中
Store reference info local variable
*/

type ASTORE struct {
	base.Index8Instruction
}

type ASTORE_0 struct {
	base.NoOperandsInstruction
}

type ASTORE_1 struct {
	base.NoOperandsInstruction
}

type ASTORE_2 struct {
	base.NoOperandsInstruction
}

type ASTORE_3 struct {
	base.NoOperandsInstruction
}

func (this *ASTORE) Execute(frame *rtda.Frame) {
	_astore(frame, this.Index)
}

func (this *ASTORE_0) Execute(frame *rtda.Frame) {
	_astore(frame, 0)
}

func (this *ASTORE_1) Execute(frame *rtda.Frame) {
	_astore(frame, 1)
}

func (this *ASTORE_2) Execute(frame *rtda.Frame) {
	_astore(frame, 2)
}

func (this *ASTORE_3) Execute(frame *rtda.Frame) {
	_astore(frame, 3)
}

func _astore(frame *rtda.Frame, index uint) {
	ref := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, ref)
}

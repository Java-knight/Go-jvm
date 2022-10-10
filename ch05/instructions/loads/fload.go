package loads

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/**
fload指令: 从局部变量表加载float到操作数栈
*/

type FLOAD struct {
	base.Index8Instruction
}

type FLOAD_0 struct {
	base.NoOperandsInstruction
}

type FLOAD_1 struct {
	base.NoOperandsInstruction
}

type FLOAD_2 struct {
	base.NoOperandsInstruction
}

type FLOAD_3 struct {
	base.NoOperandsInstruction
}

func (this *FLOAD) Execute(frame *rtda.Frame) {
	_fload(frame, this.Index)
}

func (this *FLOAD_0) Execute(frame *rtda.Frame) {
	_fload(frame, 0)
}

func (this *FLOAD_1) Execute(frame *rtda.Frame) {
	_fload(frame, 1)
}

func (this *FLOAD_2) Execute(frame *rtda.Frame) {
	_fload(frame, 2)
}

func (this *FLOAD_3) Execute(frame *rtda.Frame) {
	_fload(frame, 3)
}

func _fload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}

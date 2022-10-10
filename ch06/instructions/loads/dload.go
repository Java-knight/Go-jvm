package loads

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/**
dload指令: 从局部变量表加载double到操作数栈
*/

type DLOAD struct {
	base.Index8Instruction
}

type DLOAD_0 struct {
	base.NoOperandsInstruction
}

type DLOAD_1 struct {
	base.NoOperandsInstruction
}

type DLOAD_2 struct {
	base.NoOperandsInstruction
}

type DLOAD_3 struct {
	base.NoOperandsInstruction
}

func (this *DLOAD) Execute(frame *rtda.Frame) {
	_dload(frame, this.Index)
}

func (this *DLOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame, 0)
}

func (this *DLOAD_1) Execute(frame *rtda.Frame) {
	_dload(frame, 1)
}

func (this *DLOAD_2) Execute(frame *rtda.Frame) {
	_dload(frame, 2)
}

func (this *DLOAD_3) Execute(frame *rtda.Frame) {
	_dload(frame, 3)
}

func _dload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}

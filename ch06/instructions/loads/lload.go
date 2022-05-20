package loads

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
lload指令: 从局部变量表加载long到操作数栈
*/

type LLOAD struct {
	base.Index8Instruction
}

type LLOAD_0 struct {
	base.NoOperandsInstruction
}

type LLOAD_1 struct {
	base.NoOperandsInstruction
}

type LLOAD_2 struct {
	base.NoOperandsInstruction
}

type LLOAD_3 struct {
	base.NoOperandsInstruction
}

func (this *LLOAD) Execute(frame *rtda.Frame) {
	_lload(frame, this.Index)
}

func (this *LLOAD_0) Execute(frame *rtda.Frame) {
	_lload(frame, 0)
}

func (this *LLOAD_1) Execute(frame *rtda.Frame) {
	_lload(frame, 1)
}

func (this *LLOAD_2) Execute(frame *rtda.Frame) {
	_lload(frame, 2)
}

func (this *LLOAD_3) Execute(frame *rtda.Frame) {
	_lload(frame, 3)
}

func _lload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}

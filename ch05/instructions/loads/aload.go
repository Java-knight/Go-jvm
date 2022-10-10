package loads

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
aload指令: 从局部变量表加载引用reference到操作数栈
*/

type ALOAD struct {
	base.Index8Instruction
}

type ALOAD_0 struct {
	base.NoOperandsInstruction
}

type ALOAD_1 struct {
	base.NoOperandsInstruction
}

type ALOAD_2 struct {
	base.NoOperandsInstruction
}

type ALOAD_3 struct {
	base.NoOperandsInstruction
}

func (this *ALOAD) Execute(frame *rtda.Frame) {
	_aload(frame, this.Index)
}

func (this *ALOAD_0) Execute(frame *rtda.Frame) {
	_aload(frame, 0)
}

func (this *ALOAD_1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}

func (this *ALOAD_2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}

func (this *ALOAD_3) Execute(frame *rtda.Frame) {
	_aload(frame, 3)
}

func _aload(frame *rtda.Frame, index uint) {
	ref := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(ref)
}

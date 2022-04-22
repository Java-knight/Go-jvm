package loads

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/**
iload指令: 从局部变量表加载int到操作数栈
*/

// 将局部变量表中int类型的数据放入栈顶
type ILOAD struct {
	base.Index8Instruction
}

type ILOAD_0 struct {
	base.NoOperandsInstruction
}

type ILOAD_1 struct {
	base.NoOperandsInstruction
}

type ILOAD_2 struct {
	base.NoOperandsInstruction
}

type ILOAD_3 struct {
	base.NoOperandsInstruction
}

func (this *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, uint(this.Index))
}

func (this *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

func (this *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

func (this *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

func (this *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}

func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

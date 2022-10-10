package math

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
MUL指令: 乘指令
*/

// Multiply int
type IMUL struct {
	base.NoOperandsInstruction
}

// Multiply long
type LMUL struct {
	base.NoOperandsInstruction
}

// Multiply float
type FMUL struct {
	base.NoOperandsInstruction
}

// Multiply double
type DMUL struct {
	base.NoOperandsInstruction
}

func (this *IMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	result := val1 * val2
	stack.PushInt(result)
}

func (this *LMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopLong()
	val1 := stack.PopLong()
	result := val1 * val2
	stack.PushLong(result)
}

func (this *FMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopFloat()
	val1 := stack.PopFloat()
	result := val1 * val2
	stack.PushFloat(result)
}

func (this *DMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopDouble()
	val1 := stack.PopDouble()
	result := val1 * val2
	stack.PushDouble(result)
}

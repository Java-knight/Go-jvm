package math

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
SUB: 减法指令
*/

// Subtract int
type ISUB struct {
	base.NoOperandsInstruction
}

// Subtract long
type LSUB struct {
	base.NoOperandsInstruction
}

// Subtract float
type FSUB struct {
	base.NoOperandsInstruction
}

// Subtract double
type DSUB struct {
	base.NoOperandsInstruction
}

func (this *ISUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	result := val1 - val2
	stack.PushInt(result)
}

func (this *LSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopLong()
	val1 := stack.PopLong()
	result := val1 - val2
	stack.PushLong(result)
}

func (this *FSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopFloat()
	val1 := stack.PopFloat()
	result := val1 - val2
	stack.PushFloat(result)
}

func (this *DSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopDouble()
	val1 := stack.PopDouble()
	result := val1 - val2
	stack.PushDouble(result)
}

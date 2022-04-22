package math

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
ADD: 加法指令
*/

// add int
type IADD struct {
	base.NoOperandsInstruction
}

// add long
type LADD struct {
	base.NoOperandsInstruction
}

// add float
type FADD struct {
	base.NoOperandsInstruction
}

// add double
type DADD struct {
	base.NoOperandsInstruction
}

func (this *IADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	result := val1 + val2
	stack.PushInt(result)
}

func (this *LADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopLong()
	val1 := stack.PopLong()
	result := val1 + val2
	stack.PushLong(result)
}

func (this *FADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopFloat()
	val1 := stack.PopFloat()
	result := val1 + val2
	stack.PushFloat(result)
}

func (this *DADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopDouble()
	val1 := stack.PopDouble()
	result := val1 + val2
	stack.PushDouble(result)
}

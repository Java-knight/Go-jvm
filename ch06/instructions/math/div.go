package math

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
DIV: 除法指令
*/

// Divide int
type IDIV struct {
	base.NoOperandsInstruction
}

// Divide long
type LDIV struct {
	base.NoOperandsInstruction
}

// Divide float
type FDIV struct {
	base.NoOperandsInstruction
}

// Divide double
type DDIV struct {
	base.NoOperandsInstruction
}

func (this *IDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	result := val1 / val2
	stack.PushInt(result)
}

func (this *LDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopLong()
	val1 := stack.PopLong()
	result := val1 / val2
	stack.PushLong(result)
}

func (this *FDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopFloat()
	val1 := stack.PopFloat()
	result := val1 / val2
	stack.PushFloat(result)
}

func (this *DDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopDouble()
	val1 := stack.PopDouble()
	result := val1 / val2
	stack.PushDouble(result)
}

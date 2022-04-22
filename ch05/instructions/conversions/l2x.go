package conversions

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
L2X: long类型转其它类型int\long\double
*/

// Convert long to int
type L2I struct {
	base.NoOperandsInstruction
}

// Convert long to float
type L2F struct {
	base.NoOperandsInstruction
}

// Convert long to double
type L2D struct {
	base.NoOperandsInstruction
}

func (this *L2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	result := int32(val)
	stack.PushInt(result)
}

func (this *L2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	result := float32(val)
	stack.PushFloat(result)
}

func (this *L2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	result := float64(val)
	stack.PushDouble(result)
}

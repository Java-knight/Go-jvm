package conversions

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
F2X: float类型转其它类型int\long\double
*/

// Convert float to int
type F2I struct {
	base.NoOperandsInstruction
}

// Convert float to long
type F2L struct {
	base.NoOperandsInstruction
}

// Convert float to double
type F2D struct {
	base.NoOperandsInstruction
}

func (this *F2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	result := int32(val)
	stack.PushInt(result)
}

func (this *F2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	result := int64(val)
	stack.PushLong(result)
}

func (this *F2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	result := float64(val)
	stack.PushDouble(result)
}

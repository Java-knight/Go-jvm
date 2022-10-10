package conversions

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
I2X: int类型转其它类型byte\short\char\long\float\double
*/

// Convert int to byte
type I2B struct {
	base.NoOperandsInstruction
}

// Convert int to short
type I2S struct {
	base.NoOperandsInstruction
}

// Convert int to char
type I2C struct {
	base.NoOperandsInstruction
}

// Convert int to long
type I2L struct {
	base.NoOperandsInstruction
}

// Convert int to float
type I2F struct {
	base.NoOperandsInstruction
}

// Convert int to double
type I2D struct {
	base.NoOperandsInstruction
}

func (this *I2B) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	result := int32(int8(val))
	stack.PushInt(result)
}

func (this *I2S) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	result := int32(int16(val))
	stack.PushInt(result)
}

func (this *I2C) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	result := int32(uint16(val))
	stack.PushInt(result)
}

func (this *I2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	result := int64(val)
	stack.PushLong(result)
}

func (this *I2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	result := float32(val)
	stack.PushFloat(result)
}

func (this *I2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	result := float64(val)
	stack.PushDouble(result)
}

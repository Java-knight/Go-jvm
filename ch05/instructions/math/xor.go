package math

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
XOR: 异或指令
*/

// Boolean XOR int
type IXOR struct {
	base.NoOperandsInstruction
}

// Boolean XOR long
type LXOR struct {
	base.NoOperandsInstruction
}

func (this *IXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	result := val1 ^ val2
	stack.PushInt(result)
}

func (this *LXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopLong()
	val1 := stack.PopLong()
	result := val1 ^ val2
	stack.PushLong(result)
}

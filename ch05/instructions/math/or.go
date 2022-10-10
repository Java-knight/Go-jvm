package math

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
OR: boolean|int / boolean|long或指令
*/

// Boolean OR int
type IOR struct {
	base.NoOperandsInstruction
}

// Boolean OR long
type LOR struct {
	base.NoOperandsInstruction
}

func (this *IOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	result := val1 | val2
	stack.PushInt(result)
}

func (this *LOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopLong()
	val1 := stack.PopLong()
	result := val1 | val2
	stack.PushLong(result)
}

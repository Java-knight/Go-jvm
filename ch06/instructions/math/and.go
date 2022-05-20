package math

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
布尔运算指令: 只能操作 int 和 long 变量，分为按位与（and）、按位或（or）、按位异或（xor）3种
*/
type IAND struct {
	base.NoOperandsInstruction
}

type LAND struct {
	base.NoOperandsInstruction
}

func (this *IAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	result := val1 & val2
	stack.PushInt(result)
}

func (this *LAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopLong()
	val1 := stack.PopLong()
	result := val1 & val2
	stack.PushLong(result)
}

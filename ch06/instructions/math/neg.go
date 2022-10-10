package math

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
NEG: 求补指令
求补运算: 用零减去操作数，然后结果返回操作数。
case: index = 5   NEG index  = -5
*/

// Negate int
type INEG struct {
	base.NoOperandsInstruction
}

// Negate long
type LNEG struct {
	base.NoOperandsInstruction
}

// Negate float
type FNEG struct {
	base.NoOperandsInstruction
}

// Negate double
type DNEG struct {
	base.NoOperandsInstruction
}

func (this *INEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(-val)
}

func (this *LNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}

func (this *FNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}

func (this *DNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}

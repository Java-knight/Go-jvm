package math

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
	"math"
)

/*
求余指令
*/

// Remainder int（余数是int类型）
type IREM struct {
	base.NoOperandsInstruction
}

// Remainder long（余数是long类型）
type LREM struct {
	base.NoOperandsInstruction
}

// Remainder double（余数是double类型）
type DREM struct {
	base.NoOperandsInstruction
}

// Remainder float（余数是float类型）
type FREM struct {
	base.NoOperandsInstruction
}

func (this *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := val1 % val2
	stack.PushInt(result)
}

func (this *LREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopLong()
	val1 := stack.PopLong()
	if val2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := val1 % val2
	stack.PushLong(result)
}

func (this *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopFloat()
	val1 := stack.PopFloat()
	result := float32(math.Mod(float64(val1), float64(val2))) // 32—>64运算—>32结果
	stack.PushFloat(result)
}

func (this *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopDouble()
	val1 := stack.PopDouble()
	result := math.Mod(val1, val2)
	stack.PushDouble(result)
}

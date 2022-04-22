package conversions

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
d2x转换指令：将double类型转换成其它类型
*/

// double 转 float
type D2F struct {
	base.NoOperandsInstruction
}

// double 转 int
type D2I struct {
	base.NoOperandsInstruction
}

// double 转 long
type D2L struct {
	base.NoOperandsInstruction
}

// TODO: 三个Execute函数的逻辑基本相同，可以通过一个公共函数和一个判断函数完成
func (this *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	result := float32(val)
	stack.PushFloat(result)
}

func (this *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	result := int32(val)
	stack.PushInt(result)
}

func (this *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	result := int64(val)
	stack.PushLong(result)
}

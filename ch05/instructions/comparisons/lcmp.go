package comparisons

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
lcmp指令: 用于比较long变量
*/

type LCMP struct {
	base.NoOperandsInstruction
}

// val1 > val2  返回1
// val1 == val2 返回0
// val1 < val2  返回-1
func (this *LCMP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopLong()
	val1 := stack.PopLong()
	if val1 > val2 {
		stack.PushInt(1)
	} else if val1 == val2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}

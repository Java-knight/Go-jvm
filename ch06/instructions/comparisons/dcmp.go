package comparisons

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
dcmp指令：用于比较double类型变量，和float的逻辑基本一致
*/

type DCMPG struct {
	base.NoOperandsInstruction
}

type DCMPL struct {
	base.NoOperandsInstruction
}

func (this *DCMPG) Execute(frame *rtda.Frame) {
	_dcmp(frame, true)
}

func (this *DCMPL) Execute(frame *rtda.Frame) {
	_dcmp(frame, false)
}

// TODO if分支太多，可以合并
// 至少有一个变量NaN(Not a Number), 通过gFlag判断, dcmpg指令比较结果是1, 而dcmpl指令比较结果是-1
func _dcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	val2 := stack.PopDouble()
	val1 := stack.PopDouble()
	if val1 > val2 {
		stack.PushInt(1)
	} else if val1 == val2 {
		stack.PushInt(0)
	} else if val1 < val2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

package comparisons

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
fcmpg 和 fcmpl指令: 用于比较float变量
由于浮点数计算有可能产生NaN(Not a Number)值，所以比较两个浮点数时，除了大于、等于、小于之外还有第4种结果：无法比较。
fcmpg 和 fcmpl指令的区别就在于对第4种结果的定义
*/

type FCMPG struct {
	base.NoOperandsInstruction
}

type FCMPL struct {
	base.NoOperandsInstruction
}

func (this *FCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}

func (this *FCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}

// 编写一个公共的float比较函数
// gFlag: 两个float变量，至少有一个是NaN，用fcmpg指令比较的结果是1，而用fcmpl指令比较的结果是-1。
func _fcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	val2 := stack.PopFloat()
	val1 := stack.PopFloat()
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

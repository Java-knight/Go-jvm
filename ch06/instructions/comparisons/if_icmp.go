package comparisons

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
if_icmp<coud>指令: 把栈顶的两个int变量弹出，然后进行比较，满足条件则跳转(参考ifcond指令)
if_icmpeq: val1 == val2
if_icmpne: val1 != val2
if_icmplt: val1 < val2
if_icmple: val1 <= val2
if_icmpgt: val1 > val2
if_icmpge: val1 >= val2
*/

type If_ICMPEQ struct {
	base.BranchInstruction
}

type If_ICMPNE struct {
	base.BranchInstruction
}

type If_ICMPLT struct {
	base.BranchInstruction
}

type If_ICMPLE struct {
	base.BranchInstruction
}

type If_ICMPGT struct {
	base.BranchInstruction
}

type If_ICMPGE struct {
	base.BranchInstruction
}

func (this *If_ICMPEQ) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 == val2 {
		base.Branch(frame, this.Offset)
	}
}

func (this *If_ICMPNE) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 != val2 {
		base.Branch(frame, this.Offset)
	}
}

func (this *If_ICMPLT) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 < val2 {
		base.Branch(frame, this.Offset)
	}
}

func (this *If_ICMPLE) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 <= val2 {
		base.Branch(frame, this.Offset)
	}
}

func (this *If_ICMPGT) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 > val2 {
		base.Branch(frame, this.Offset)
	}
}

func (this *If_ICMPGE) Execute(frame *rtda.Frame) {
	if val1, val2 := _icmpPop(frame); val1 >= val2 {
		base.Branch(frame, this.Offset)
	}
}

// 出栈val1 和 val2
func _icmpPop(frame *rtda.Frame) (val1, val2 int32) {
	stack := frame.OperandStack()
	val2 = stack.PopInt()
	val1 = stack.PopInt()
	return
}

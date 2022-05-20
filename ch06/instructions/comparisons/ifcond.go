package comparisons

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
if<cond>指令: 把操作数栈顶的int变量弹出, 然后跟0进行比较, 满足条件则跳转
假设从栈顶弹出的变量是x, 则指令执行跳转操作条件如下:
ifeq: x == 0
ifne: x != 0
iflt: x < 0
ifle: x <= 0
ifgt: x > 0
ifge: x >= 0
*/

type IFEQ struct {
	base.BranchInstruction
}

type IFNE struct {
	base.BranchInstruction
}

type IFLT struct {
	base.BranchInstruction
}

type IFLE struct {
	base.BranchInstruction
}

type IFGT struct {
	base.BranchInstruction
}

type IFGE struct {
	base.BranchInstruction
}

func (this *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, this.Offset)
	}
}

func (this *IFNE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, this.Offset)
	}
}

func (this *IFLT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, this.Offset)
	}
}

func (this *IFLE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, this.Offset)
	}
}

func (this *IFGT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, this.Offset)
	}
}

func (this *IFGE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, this.Offset)
	}
}

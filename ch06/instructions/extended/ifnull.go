package extended

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
ifnull 和 ifnonnull指令: 根据引用是否 是null进行跳转
ifnull指令: if分支引用是null，走ifnull指令
ifnonnull指令: if分支引用不是null，走ifnonnull指令
*/

type IFNULL struct {
	base.BranchInstruction
}

type IFNONNULL struct {
	base.BranchInstruction
}

func (this *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, this.Offset)
	}
}

func (this *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, this.Offset)
	}
}

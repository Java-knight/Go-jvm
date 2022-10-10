package control

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
goto指令: 进行无条件跳转
*/

type GOTO struct {
	base.BranchInstruction
}

func (this *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, this.Offset)
}

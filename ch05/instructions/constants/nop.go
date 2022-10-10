package constants

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

// nop指令：没有什么作用，Do nothing
type NOP struct {
	base.NoOperandsInstruction
}

// 执行
func (this *NOP) Execute(frame *rtda.Frame) {
	// 什么也不做
}

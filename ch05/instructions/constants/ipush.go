package constants

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/**
bipush指令：从操作数中获取一个byte型整数，扩展成int类型，放入栈顶
sipush指令：从操作数中获取一个short型整数，扩展成int类型，放入栈顶
*/

type BIPUSH struct {
	val int8
}

type SIPUSH struct {
	val int16
}

func (this *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	this.val = reader.ReadInt8()
}

func (this *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(this.val)
	frame.OperandStack().PushInt(i)
}

func (this *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	this.val = reader.ReadInt16()
}

func (this *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(this.val)
	frame.OperandStack().PushInt(i)
}

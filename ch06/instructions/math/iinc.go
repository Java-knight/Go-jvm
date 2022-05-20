package math

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
iinc指令: 给局部变量表中的int变量增加常量值, 局部变量表索引和常量值都由指令的操作数提供
*/

type IINC struct {
	Index uint  // 局部变量表索引
	Const int32 // 常量值
}

// 从字节码读取操作数
func (this *IINC) FetchOperands(reader *base.BytecodeReader) {
	this.Index = uint(reader.ReadUint8())
	this.Const = int32(reader.ReadInt32())
}

// 从局部变量表中读取变量，给它加上常量值，再把结果写回局部变量表
func (this *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(this.Index)
	val += this.Const
	localVars.SetInt(this.Index, val)
}

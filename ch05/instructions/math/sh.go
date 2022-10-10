package math

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
位移指令：分为左移（int\long） 和 右移[算术int\long右移; 逻辑int\long右移]
*/

// Shift left int（左移int）
type ISHL struct {
	base.NoOperandsInstruction
}

// Arithmetic Shift right int（算术右移int）【算术就是有符号】
type ISHR struct {
	base.NoOperandsInstruction
}

// Logical Shift right int（逻辑右移int）【逻辑就是无符号】
type IUSHR struct {
	base.NoOperandsInstruction
}

// Shift left int（左移long）
type LSHL struct {
	base.NoOperandsInstruction
}

// Arithmetic Shift right long（算术右移long）【算术就是有符号】
type LSHR struct {
	base.NoOperandsInstruction
}

// Logical Shift right long（逻辑右移long）【逻辑就是无符号】
type LUSHR struct {
	base.NoOperandsInstruction
}

// val1 << int(val2)
// int类型是32位，最大就是5个比特，0x1f就是5个比特"0001 1111"（f就是1111）
// Go语言在进行位运算必须是无符号整数
func (this *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	s := uint32(val2) & 0x1f
	result := val1 << s
	stack.PushInt(result)
}

func (this *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	s := uint32(val2) & 0x1f
	result := val1 >> s
	stack.PushInt(result)
}

func (this *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	s := uint32(val2) & 0x1f
	result := int32(uint32(val1) >> s)
	stack.PushInt(result)
}

// long变量有64位，所以取val2的前6个比特，0x3f就是"0011 1111"（3就是0011，f就是1111）
func (this *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopLong()
	s := uint32(val2) & 0x3f
	result := val1 << s
	stack.PushLong(result)
}

// long变量有64位，所以取val2的前6个比特，0x3f就是"0011 1111"（3就是0011，f就是1111）
func (this *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopLong()
	s := uint32(val2) & 0x3f
	result := val1 >> s
	stack.PushLong(result)
}

func (this *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopLong()
	s := uint32(val2) & 0x3f
	result := int32(uint32(val1) >> s)
	stack.PushInt(result)
}

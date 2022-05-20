package constants

import (
	"Go-jvm/ch06/instructions/base"
	"Go-jvm/ch06/rtda"
)

/*
ldc系列指令：从运行时常量池中加载常量值，并把它推入操作数栈
ldc指令：加载 int、float
ldc_w指令：字符串常量
ldc2_w指令：加载long 和 double常量
*/

type LDC struct {
	base.Index8Instruction
}

type LDC_W struct {
	base.Index16Instruction
}

type LDC2_W struct {
	base.Index16Instruction
}

func (this *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame, this.Index)
}

func (this *LDC_W) Execute(frame *rtda.Frame) {
	_ldc(frame, this.Index)
}

func _ldc(frame *rtda.Frame, index uint) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(index)
	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	// TODO case string
	// TODO case *heap.ClassRef
	default:
		panic("todo: ldc!")
	}
}

func (this *LDC2_W) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(this.Index)
	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		panic("java.lang.ClassFor")
	}
}

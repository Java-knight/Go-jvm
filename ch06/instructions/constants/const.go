package constants

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
  将隐含栈操作码中的常量值放入操作数栈的栈顶
*/

// 常量空值 aconst_null，作用：push null
type ACONST_NULL struct {
	base.NoOperandsInstruction
}

// push double
type DCONST_0 struct {
	base.NoOperandsInstruction
}
type DCONST_1 struct {
	base.NoOperandsInstruction
}

// push float
type FCONST_0 struct {
	base.NoOperandsInstruction
}
type FCONST_1 struct {
	base.NoOperandsInstruction
}
type FCONST_2 struct {
	base.NoOperandsInstruction
}

// push int
type ICONST_M1 struct {
	base.NoOperandsInstruction
}
type ICONST_0 struct {
	base.NoOperandsInstruction
}
type ICONST_1 struct {
	base.NoOperandsInstruction
}
type ICONST_2 struct {
	base.NoOperandsInstruction
}
type ICONST_3 struct {
	base.NoOperandsInstruction
}
type ICONST_4 struct {
	base.NoOperandsInstruction
}
type ICONST_5 struct {
	base.NoOperandsInstruction
}

// push long
type LCONST_0 struct {
	base.NoOperandsInstruction
}
type LCONST_1 struct {
	base.NoOperandsInstruction
}

func (this *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

func (this *DCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
}
func (this *DCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

func (this *FCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
}
func (this *FCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(1.0)
}
func (this *FCONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(2.0)
}

func (this *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}
func (this *ICONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}
func (this *ICONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}
func (this *ICONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(2)
}
func (this *ICONST_3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(3)
}
func (this *ICONST_4) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(4)
}
func (this *ICONST_5) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(5)
}

func (this *LCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}
func (this *LCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}

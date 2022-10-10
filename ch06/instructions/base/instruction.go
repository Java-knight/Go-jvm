package base

import "Go-jvm/ch05/rtda"

// 字节码指令接口（相当于Java中的抽象类）
type Instruction interface {
	FetchOperands(reader *BytecodeReader) // 按照操作数类型实现一些结构体, 并实现FetchOperands方法（获取操作数）
	Execute(frame *rtda.Frame)            // 执行，使用栈桢
}

// 非操作数指令
type NoOperandsInstruction struct {
}

// 跳转指令
type BranchInstruction struct {
	Offset int // 偏移量
}

// index8指令: 局部变量表index索引是由单个字节码操作数（存储和加载类指令需要根据索引存取）
type Index8Instruction struct {
	Index uint
}

// index16指令: 获取运行时常量池索引是由两个字节操作数
type Index16Instruction struct {
	Index uint
}

func (this *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}

func (this *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	this.Offset = int(reader.ReadInt16())
}

func (this *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	this.Index = uint(reader.ReadUint8())
}

func (this *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	this.Index = uint(reader.ReadUint16())
}

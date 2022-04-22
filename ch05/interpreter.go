package main

import (
	"Go-jvm/ch05/classfile"
	"Go-jvm/ch05/instructions"
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
	"fmt"
)

/*
这个只是一个简单版本的解释器，为了dev测试当前功能
解释器: JVM的执行Java方法的，Java虚拟机规范必备的一个组件
*/

func interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()
	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)
	defer catchErr(frame) // 因为解释器还没有实现return的指令，可以先用catchErr()函数处理
	loop(thread, bytecode)
}

// 处理未实现return指令的异常
func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars: %v\n", frame.LocalVars())
		fmt.Printf("OPerandStack: %v\n", frame.OperandStack())
		panic(r)
	}
}

// 把局部变量表和操作数栈中内容打印出来(计算pc、解码指令、执行指令)
func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		pc := frame.NextPC()
		thread.SetPC(pc)
		// 解码decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		instruction := instructions.NewInstruction(opcode)
		instruction.FetchOperands(reader) // 获取操作数
		frame.SetNextPC(reader.PC())
		// 执行execute
		fmt.Printf("PC: %2d instruction: %T %v\n", pc, instruction, instruction)
		instruction.Execute(frame)
	}
}

package references

import (
	"Go-jvm/ch06/instructions/base"
	"Go-jvm/ch06/rtda"
	"Go-jvm/ch06/rtda/heap"
)

/*
getstatic指令：获取类的某个静态变量值，然后推入栈顶
只需要一个操作数：uint16 常量池索引
*/

// Get static field from class
type GET_STATIC struct {
	base.Index16Instruction
}

func (this *GET_STATIC) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	fieldRef := cp.GetConstant(this.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()

	// 如果解析后字段不是静态字段，也要抛IncompatibleClassChangeError异常。
	// 如果声明字段的类还没有初始化好，也需要先初始化。
	// getstatic 只是读取静态变量的值，自然也就不用管它是否 final 了
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()

	// 根据字段类型，从静态变量中取出来对应的值，推入操作数栈 栈顶
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))
	}
}

package references

import (
	"Go-jvm/ch06/instructions/base"
	"Go-jvm/ch06/rtda"
	"Go-jvm/ch06/rtda/heap"
)

/*
getfield指令：获取对象的实例变量值，然后推入操作数栈，它需要两个操作数。
第一个操作数是 uint16 常量池索引；第二个操作数是对象引用
*/

// Fetch field from object
type GET_FIELD struct {
	base.Index16Instruction
}

func (this *GET_FIELD) Execute(frame *rtda.Frame) {
	// (1) 字段符号解析
	cp := frame.Method().Class().ConstantPool()
	fieldRef := cp.GetConstant(this.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// (2) 弹出对象如果是nil，会抛NPE
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := ref.Fields()
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

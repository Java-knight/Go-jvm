package references

import (
	"Go-jvm/ch06/instructions/base"
	"Go-jvm/ch06/rtda"
	"Go-jvm/ch06/rtda/heap"
)

/*
putfield指令：给实例变量赋值，它需要三个操作数。
前两个操作数是常量池索引和变量值，用法和putstatic一样。第三个操作符是对象引用，从操作数栈弹出
*/

// Set field in Object
type PUT_FIELD struct {
	base.Index16Instruction
}

func (this *PUT_FIELD) Execute(frame *rtda.Frame) {
	// (1) 先拿到当前方法、当前类和当前常量池，然后解析字段符号引用。
	// 如果声明字段的类还没有被初始化，则需要先初始化该类
	curMethod := frame.Method()
	curClass := curMethod.Class()
	cp := curClass.ConstantPool()
	fieldRef := cp.GetConstant(this.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()

	// (2) 解析后的字段必须是实例字段，否则抛出 IncompatibleClassChangeError
	// 如果是 final 字段，则只能在构造函数中初始化，否则抛出IllegalAccessError
	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if field.IsFinal() {
		if curClass != field.Class() || curMethod.Name() != "<init>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	// (3) 先根据字段类型从操作数栈中弹出相应的变量，然后弹出对象引用。如果是null，需要抛出NPE，否则通过引用给实例变量赋值
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	stack := frame.OperandStack()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		val := stack.PopInt()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetInt(slotId, val)
	case 'F':
		val := stack.PopFloat()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetFloat(slotId, val)
	case 'J':
		val := stack.PopLong()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetLong(slotId, val)
	case 'D':
		val := stack.PopDouble()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetDouble(slotId, val)
	case 'L', '[':
		val := stack.PopRef()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetRef(slotId, val)
	default:
		// TODO
	}
}

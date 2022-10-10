package references

import (
	"Go-jvm/ch06/instructions/base"
	"Go-jvm/ch06/rtda"
	"Go-jvm/ch06/rtda/heap"
)

/*
putstatic指令：给类的某个静态变量赋值，它需要两个操作数
（1）第一个操作数是 uint16索引，来自字节码。
通过这个索引可以从当前类的运行时常量池中找到一个字段符号引用，解析这个符号就可以知道要给类的哪个静态变量赋值。
（2）第二个操作数是要赋给静态变量的值，从操作数栈中弹出
*/

// Set static field in class
type PUT_STATIC struct {
	base.Index16Instruction
}

func (this *PUT_STATIC) Execute(frame *rtda.Frame) {
	// (1) 先拿到当前方法、当前类和当前常量池，然后解析字段符号引用。
	// 如果声明字段的类还没有被初始化，则需要先初始化该类
	curMethod := frame.Method()
	curClass := curMethod.Class()
	cp := curClass.ConstantPool()
	fieldRef := cp.GetConstant(this.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()

	// (2) 如果解析后中字段是实例字段而非静态字段，则抛出IncompatibleClassChangeError异常
	// 如果是final 字段，则实际操作的是静态常量，只能在类初始化方法中给它赋值。否则会抛IllegalAccessError异常
	// 注：类的初始化方法由编译器生成，名字<clinit>
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if field.IsFinal() {
		if curClass != class || curMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	// (3) 获取到字段的描述符descriptor[0]，根据字段类型从操作数栈弹出相应的值，然后赋给静态变量
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		slots.SetInt(slotId, stack.PopInt())
	case 'F':
		slots.SetFloat(slotId, stack.PopFloat())
	case 'J':
		slots.SetLong(slotId, stack.PopLong())
	case 'D':
		slots.SetDouble(slotId, stack.PopDouble())
	case 'L', '[':
		slots.SetRef(slotId, stack.PopRef())
	default:
		// TODO
	}
}

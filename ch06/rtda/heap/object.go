package heap

// 目前就是一个临时的结构体
type Object struct {
	class *Class // 对象 Class 指针
	field Slots
}

// 创建对象(这里实例变量已经赋好值类，不需要额外的操作)
func newObject(class *Class) *Object {
	return &Object{
		class: class,
		field: newSlots(class.instanceSlotCount),
	}
}

// getter class/field
func (this *Object) Class() *Class {
	return this.class
}
func (this *Object) Fields() Slots {
	return this.field
}

// 如果对象不是 null，则解析类符号引用，判断对象是否是类的实例，然后把判断解决推入操作数栈。
// Java虚拟机规范给出类具体的操作步骤：就是Object.IsInstanceOf()
func (this *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(this.class)
}

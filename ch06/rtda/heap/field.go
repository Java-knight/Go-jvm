package heap

import (
	"Go-jvm/ch06/classfile"
)

/*
成员变量

问题：如何直到静态变量和实例变量需要多少空间？以及哪个字段对应 Slots 中的哪个位置？
（1）第一个问题，只要数一下类的字段即可。假设某个类有 m 个静态字段和 n 个实例字段，那么静态变量和实例变量所需的空间大小就分别是 m' 和 n'。
这里需要注意两点: 首先，类是可以继承的，也就是说，在统计实例变量个数时，要递归地数超类的实例变量；其次，long 和 double 字段都占两个slot位置，所以 m' >= m，n' >= n
（2）第二个问题，在统计字段时，给字段按编号上号就可以了。
这里需要注意三点: 首先，静态字段和实例字段要分开编号，否则会混乱；其次，对于实例字段，一定要从继承关系的最顶端，也就是 java.lang.Object 开始编号，否则也会混乱；
               最后，编号时也是考虑 long 和 double 类型。
*/

type Field struct {
	ClassMember
	slotId          uint // slot的索引，key
	constValueIndex uint // 常量池索引，value
}

// 创建类所有的 成员变量
func newFields(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
		fields[i].copyAttributes(cfField)
	}
	return fields
}

// 属性拷贝(setter 属性index)
func (this *Field) copyAttributes(cfField *classfile.MemberInfo) {
	if valAttr := cfField.ConstantValueAttribute(); valAttr != nil {
		this.constValueIndex = uint(valAttr.ConstantValueIndex())
	}
}

// check vars(volatile\transient\enum)
func (this *Field) IsVolatile() bool {
	return 0 != this.accessFlags&ACC_VOLATILE
}
func (this *Field) IsTransient() bool {
	return 0 != this.accessFlags&ACC_TRANSIENT
}
func (this *Field) IsEnum() bool {
	return 0 != this.accessFlags&ACC_ENUM
}
func (this *Field) isLongOrDouble() bool { // 判断是否是long\double
	return this.descriptor == "J" || this.descriptor == "D"
}

// getter vars(slotId\constValueIndex)
func (this *Field) ConstValueIndex() uint {
	return this.constValueIndex
}
func (this *Field) SlotId() uint {
	return this.slotId
}

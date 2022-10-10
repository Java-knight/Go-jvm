package heap

import "math"

// ｜-----｜-----｜-----｜-----｜
// ｜slot1｜     ｜     ｜     ｜
// ｜--- -｜-----｜-----｜-----｜
// ｜slot2｜     ｜     ｜     ｜
// ｜-----｜-----｜-----｜-----｜

/*
本来这个slot是存在于虚拟机栈的栈桢中的局部变量表的slot，这个时候需要就需要给它填入byte\short\int\float\ long\double
但是rtda包已经依赖类heap包，heap包现在不能在依赖rtda包，这个是否就会产生一些重复代码，把slot.go和local_vars.go中的代码放一份中heap包中
*/

// 局部变量表中的Slot槽位, 因为局部变量表需要用数组来表示（使用slot就可以看作数组的一个元素，方便操作）
type Slot struct {
	num int32   // 基本数据类型（byte\short\char\int\long\float\double\bool）
	ref *Object // 引用(接口类型、类类型、数组类型、null类型)
}

type Slots []Slot

// 创建所有的Slot需要的内存空间
func newSlots(slotCount uint) Slots {
	if slotCount > 0 {
		return make([]Slot, slotCount)
	}
	return nil
}

// setter/getter int
func (this Slots) SetInt(index uint, val int32) {
	this[index].num = val
}
func (this Slots) GetInt(index uint) int32 {
	return this[index].num
}

// setter/getter long(two slots)
func (this Slots) SetLong(index uint, val int64) {
	this[index].num = int32(val)
	this[index+1].num = int32(val >> 32)
}
func (this Slots) GetLong(index uint) int64 {
	low := uint32(this[index].num)
	high := uint32(this[index+1].num)
	return int64(high)<<32 | int64(low)
}

// setter/getter float
func (this Slots) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	this[index].num = int32(bits)
}
func (this Slots) GetFloat(index uint) float32 {
	bits := uint32(this[index].num)
	return math.Float32frombits(bits)
}

// setter/getter double
func (this Slots) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	this.SetLong(index, int64(bits))
}
func (this Slots) GetDouble(index uint) float64 {
	bits := uint64(this.GetLong(index))
	return math.Float64frombits(bits)
}

// setter/getter int
func (this Slots) SetRef(index uint, ref *Object) {
	this[index].ref = ref
}
func (this Slots) GetRef(index uint) *Object {
	return this[index].ref
}

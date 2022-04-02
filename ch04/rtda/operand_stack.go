package rtda

import "math"

// 操作数栈
type OperandStack struct {
	size  uint   // operand stack大小（用于记录）
	slots []Slot // 切片实现栈（数组栈）
}

// 创建一个操作数栈
func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

// 入栈 一个int类型（byte\short\char\int\boolean）
func (this *OperandStack) PushInt(val int32) {
	this.slots[this.size].num = val
	this.size++
}

// 出栈 一个int类型
func (this *OperandStack) PopInt() int32 {
	this.size--
	return this.slots[this.size].num
}

// 入栈一个 float类型（float -> uint32 -> int32）
func (this *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	this.slots[this.size].num = int32(bits)
	this.size++
}

// 出栈一个 float类型（int32 -> uint32 -> float）
func (this *OperandStack) PopFloat() float32 {
	this.size--
	bits := uint32(this.slots[this.size].num)
	return float32(math.Float32frombits(bits))
}

// 入栈一个 long类型（需要将long类型拆分成两个int，高32位 和 低32位分开放）
func (this *OperandStack) PushLong(val int64) {
	this.slots[this.size].num = int32(val)
	this.slots[this.size+1].num = int32(val >> 32)
	this.size += 2
}

// 出栈一个 long类型（需要将index和index+1位置的数据合起来, index存放中是数据的低32位，index+1存放的是数据的高32位）
func (this *OperandStack) PopLong() int64 {
	this.size -= 2
	low := uint32(this.slots[this.size].num)    // 低32位
	high := uint32(this.slots[this.size+1].num) // 高32位
	return int64(high)<<32 | int64(low)
}

// 入栈一个 double类型（原理 long+float）
func (this *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	this.PushLong(int64(bits))
}

// 出栈一个 double类型
func (this *OperandStack) PopDouble() float64 {
	bits := uint64(this.PopLong())
	return math.Float64frombits(bits)
}

// 入栈一个 引用类型
func (this *OperandStack) PushRef(ref *Object) {
	this.slots[this.size].ref = ref
	this.size++
}

// 出栈一个 引用类型（因为是切片实现的栈，在删除引用类型后，需要将内存空间指向nil）
func (this *OperandStack) PopRef() *Object {
	this.size--
	ref := this.slots[this.size].ref
	this.slots[this.size].ref = nil
	return ref
}

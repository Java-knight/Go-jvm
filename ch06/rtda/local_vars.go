package rtda

import (
	"Go-jvm/ch06/rtda/heap"
	"math"
)

// 局部变量表
type LocalVars []Slot

// 创建一个局部变量表
func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

// setting 一个int类型局部变量（slot）
func (this LocalVars) SetInt(index uint, val int32) {
	this[index].num = val
}

// getting 一个int类型局部变量（byte\short\char\boolean都是存放中int类型的）
func (this LocalVars) GetInt(index uint) int32 {
	return this[index].num
}

// setting float类型的局部变量（float -> uint32 -> int32）
func (this LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	this[index].num = int32(bits)
}

// getting float类型的局部变量（int32 -> uint32 -> float）
func (this LocalVars) GetFloat(index uint) float32 {
	bits := uint32(this[index].num)
	return math.Float32frombits(bits)
}

// setting long类型的局部变量（需要将long类型拆分成两个int，高32位 和 低32位分开放）
func (this LocalVars) SetLong(index uint, val int64) {
	this[index].num = int32(val)         // 低32位
	this[index+1].num = int32(val >> 32) // 高32位
}

// getting long类型局部变量（需要将index和index+1位置的数据合起来, index存放中是数据的低32位，index+1存放的是数据的高32位）
func (this LocalVars) GetLong(index uint) int64 {
	low := uint32(this[index].num)
	high := uint32(this[index+1].num)
	return int64(high)<<32 | int64(low)
}

// setting double类型的局部变量（原理同 long类型+float类型）
func (this LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	this.SetLong(index, int64(bits))
}

// getting double类型的局部变量
func (this LocalVars) GetDouble(index uint) float64 {
	bits := uint64(this.GetLong(index))
	return math.Float64frombits(bits)
}

// setting 引用类型的局部变量
func (this LocalVars) SetRef(index uint, ref *heap.Object) {
	this[index].ref = ref
}

// getting 引用类型的局部变量
func (this LocalVars) GetRef(index uint) *heap.Object {
	return this[index].ref
}

package classfile

import "math"

// Integer常量信息（byte\short\boolean\char类型都是存放在这个里面）
type ConstantIntegerInfo struct {
	val int32
}

// Float常量信息
type ConstantFloatInfo struct {
	val float32
}

// Long常量信息
type ConstantLongInfo struct {
	val int64
}

// Double常量信息
type ConstantDoubleInfo struct {
	val float64
}

func (this *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	this.val = int32(bytes)
}

func (this *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	this.val = math.Float32frombits(bytes) // 将bytes转换为float32
}

func (this *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	this.val = int64(bytes)
}

func (this *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	this.val = math.Float64frombits(bytes) // 将bytes转换为float64
}

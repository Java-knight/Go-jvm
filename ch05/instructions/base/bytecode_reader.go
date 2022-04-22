package base

// 字节码读取
type BytecodeReader struct {
	code []byte // 字节码
	pc   int    // 读到哪个字节
}

// 设置字节码，避免遇到重复的字节码会再次创建
func (this *BytecodeReader) Reset(code []byte, pc int) {
	this.code = code
	this.pc = pc
}

// getting pc
func (this *BytecodeReader) PC() int {
	return this.pc
}

// 读取int16字节的（short）
func (this *BytecodeReader) ReadInt8() int8 {
	return int8(this.ReadUint8())
}

// 读取uint8(byte)【读取局部变量的index...】
func (this *BytecodeReader) ReadUint8() uint8 {
	index := this.code[this.pc]
	this.pc++
	return index
}

// 读取int16字节的（short）
func (this *BytecodeReader) ReadInt16() int16 {
	return int16(this.ReadUint16())
}

// 读取uint16(short)【读取运行时常量池的index...】
func (this *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(this.ReadUint8()) // 高8位
	byte2 := uint16(this.ReadUint8()) // 低8位
	return (byte1 << 8) | byte2
}

// 读取int32（int）(byte1 byte2 byte3 byte4)
func (this *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(this.ReadInt8())
	byte2 := int32(this.ReadInt8())
	byte3 := int32(this.ReadInt8())
	byte4 := int32(this.ReadInt8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

// tableswitch指令操作码后面有0～3字节的padding, 以保证defaultOffset在字节码中的地址是4的倍数
func (this *BytecodeReader) SkipPadding() {
	for this.pc%4 != 0 {
		this.ReadUint8()
	}
}

// tableswitch 和 lookupswitch指令批量读取case跳转所需字节码的偏移量
func (this *BytecodeReader) ReadInt32s(size int32) []int32 {
	result := make([]int32, size)
	for i := range result {
		result[i] = this.ReadInt32()
	}
	return result
}

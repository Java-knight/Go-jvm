package classfile

import "encoding/binary"

// ClassReader 只是[]byte类型的包装, 读取处各个数据类型（u1、u2、u4）
// 注意：ClassReader 并没有使用索引记录数据位置, 而是使用Go语言的 reslice（切片）记录
type ClassReader struct {
	data []byte
}

// 解析class字节流中的uint8(u1)
func (this *ClassReader) readUint8() uint8 {
	val := this.data[0]
	this.data = this.data[1:]
	return val
}

// 解析class字节流中的uint16(u2)
// 返回的是class字节流数据的大端uint16
func (this *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(this.data)
	this.data = this.data[2:]
	return val
}

// 解析class字节流中的uint32(u4)
// 返回的是class字节流数据的大端uint32(大端就是高32位)
func (this *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(this.data)
	this.data = this.data[4:]
	return val
}

// 解析class字节流中的uint64( Java虚拟机规范 并没有定义u8)
func (this *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(this.data)
	this.data = this.data[8:]
	return val
}

// 读取uint16表，表的大小由开头的uint16数据指出
func (this *ClassReader) readUint16s() []uint16 {
	n := this.readUint16()
	res := make([]uint16, n)
	for i := range res {
		res[i] = this.readUint16()
	}
	return res
}

// 用于读取指定数量的字节
func (this *ClassReader) readByte(n uint32) []byte {
	bytes := this.data[:n]
	this.data = this.data[n:]
	return bytes
}

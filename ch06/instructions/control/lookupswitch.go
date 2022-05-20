package control

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
lookup_switch指令: switch-case参考tableswitch指令的注释
*/

type LOOKUP_SWITCH struct {
	defaultOffset int32   // switch-case的default语句的字节码偏移量
	npairs        int32   // switch-case的case语句数量
	matchOffsets  []int32 // 类似与Map, key就是case, value就是跳转偏移量
}

func (this *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	this.defaultOffset = reader.ReadInt32()
	this.npairs = reader.ReadInt32()
	this.matchOffsets = reader.ReadInt32s(this.npairs << 1)
}

func (this *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	key := frame.OperandStack().PopInt()
	len := this.npairs << 1
	for i := int32(0); i < len; i += 2 {
		if this.matchOffsets[i] == key {
			offset := this.matchOffsets[i+1]
			base.Branch(frame, int(offset)) // case
			return
		}
	}
	base.Branch(frame, int(this.defaultOffset)) // default
}

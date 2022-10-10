package control

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
tableswitch指令: 操作数比较复杂。
  Java语言中的switch-case语句有两种实现方式:
(1) 如果case值可以编码成一个索引表, 则实现成tableswitch指令
(2) 否则实现lookupswitch指令

int chooseNear(int i) {
	switch (i) {
	case 0:
		return 0;
	case 1:
		return 1;
	case 2:
		return 2;
	default:
		return -1;
	}
}
*/

type TABLE_SWITCH struct {
	defaultOffset int32   // 默认情况下执行跳转所需的字节码偏移量，switch-case的default语句的字节码偏移量
	low           int32   // case的取值范围（下限）
	high          int32   // case的取值范围（上限）
	jumpOffsets   []int32 // 索引表，里面存放high-low+1个case的跳转所需的字节码偏移量
}

// 获取到字节码
func (this *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	this.defaultOffset = reader.ReadInt32()
	this.low = reader.ReadInt32()
	this.high = reader.ReadInt32()
	jumpOffsetsCount := this.high - this.low + 1
	this.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (this *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= this.low && index <= this.high { // case的偏移
		offset = int(this.jumpOffsets[index-this.low])
	} else { // default的偏移
		offset = int(this.defaultOffset)
	}
	base.Branch(frame, offset)
}

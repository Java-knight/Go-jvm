package extended

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
goto_w指令: 和goto指令的唯一区别就是索引从2个字节变成了4个字节
*/

type GOTO_W struct {
	offset int
}

func (this *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	this.offset = int(reader.ReadInt32())
}

func (this *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, this.offset)
}

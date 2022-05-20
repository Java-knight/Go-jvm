package extended

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/instructions/loads"
	"Go-jvm/ch05/instructions/math"
	"Go-jvm/ch05/instructions/stores"
	"Go-jvm/ch05/rtda"
)

/*
加载类指令、存储类指令、ret 指令和 iinc 指令需要按索引访问局部变量表，索引以uint8的形式存在字节码中。
对于大部分方法来说，局部变量表不会超过256，所以一个字节表示索引就够来。但是如果有方法的局部变量表超过来这个限制呢？
Java虚拟机规范定义来 wide 指令来扩展前述指令

wide指令：改变其它指令的行为，modifiedInstruction字段存放被改变的指令。wide指令需要自己解码出modifiedInstruction
*/

type WIDE struct {
	modifiedInstruction base.Instruction // 存放被改变的指令
}

func (this *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15: // iload
		inst := &loads.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		this.modifiedInstruction = inst
	case 0x16: // lload
		inst := &loads.LLOAD{}
		inst.Index = uint(reader.ReadUint16())
		this.modifiedInstruction = inst
	case 0x17: // fload
		inst := &loads.FLOAD{}
		inst.Index = uint(reader.ReadUint16())
		this.modifiedInstruction = inst
	case 0x18: // dload
		inst := &loads.DLOAD{}
		inst.Index = uint(reader.ReadUint16())
		this.modifiedInstruction = inst
	case 0x19: // aload
		inst := &loads.ALOAD{}
		inst.Index = uint(reader.ReadUint16())
		this.modifiedInstruction = inst
	case 0x36: // istore
		inst := &stores.ISTORE{}
		inst.Index = uint(reader.ReadUint16())
		this.modifiedInstruction = inst
	case 0x37: // lstore
		inst := &stores.LSTORE{}
		inst.Index = uint(reader.ReadUint16())
		this.modifiedInstruction = inst
	case 0x38: // fstore
		inst := &stores.FSTORE{}
		inst.Index = uint(reader.ReadUint16())
		this.modifiedInstruction = inst
	case 0x39: // dstore
		inst := &stores.DSTORE{}
		inst.Index = uint(reader.ReadUint16())
		this.modifiedInstruction = inst
	case 0x3a: // astore
		inst := &stores.ASTORE{}
		inst.Index = uint(reader.ReadUint16())
		this.modifiedInstruction = inst
	case 0x84: // iinc
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadInt16())
		this.modifiedInstruction = inst
	case 0xa9: // retl: return指令没有实现，暂时panic
		panic("Unsupported opcode: 0xa9!")
	}
}

// 这个wide指令是不干活的，将指令扩展后返回，交给对应的指令去干活
func (this *WIDE) Execute(frame *rtda.Frame) {
	this.modifiedInstruction.Execute(frame)
}

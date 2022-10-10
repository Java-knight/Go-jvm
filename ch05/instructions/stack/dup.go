package stack

import (
	"Go-jvm/ch05/instructions/base"
	"Go-jvm/ch05/rtda"
)

/*
dup: 复制指令
*/

// 复制操作数栈顶顶值并插入两次栈顶值
type DUP struct {
	base.NoOperandsInstruction
}

// 复制操作数栈顶的值并向下插入两个值
type DUP_X1 struct {
	base.NoOperandsInstruction
}

type DUP_X2 struct {
	base.NoOperandsInstruction
}

type DUP2 struct {
	base.NoOperandsInstruction
}

type DUP2_X1 struct {
	base.NoOperandsInstruction
}

type DUP2_X2 struct {
	base.NoOperandsInstruction
}

/*
        ——> | d |
| d | ——|   | d |
| c |       | c |
| b |       | b |
|_a_|       |_a_|
*/
// 复制栈顶单个变量（取出栈顶变量，在让入栈顶2次）
// TODO 放入栈顶两次，可以优化，给OperandStack提供一个peek方法，只需要入栈一次
func (this *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

/*
            | d |
| d | ——|   | c |
| c |    —> | d |
| b |       | b |
|_a_|       |_a_|
*/
func (this *DUP_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot() // d
	slot2 := stack.PopSlot() // c
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

/*
            | d |
| d | ——    | c |
| c |   |   | b |
| b |    ——>| d |
|_a_|       |_a_|
*/
func (this *DUP_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot() // d
	slot2 := stack.PopSlot() // c
	slot3 := stack.PopSlot() // b
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

/*
           top
[...][c][b][a]____
          \____   |
               |  |
               V  V
[...][c][b][a][b][a]
*/
func (this *DUP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot() // a
	slot2 := stack.PopSlot() // b
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

/*
           top
[...][c][b][a]
       _/ __/
      |  |
      V  V
[...][b][a][c][b][a]
*/
func (this *DUP2_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot() // a
	slot2 := stack.PopSlot() // b
	slot3 := stack.PopSlot() // c
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

/*
              top
[...][d][c][b][a]
       ____/ __/
      |   __/
      V  V
[...][b][a][d][c][b][a]
*/
func (this *DUP2_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot() // a
	slot2 := stack.PopSlot() // b
	slot3 := stack.PopSlot() // c
	slot4 := stack.PopSlot() // d
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot4)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

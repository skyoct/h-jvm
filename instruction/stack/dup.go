package stack

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

//dup总共有6个指令，分别是dup、dup_x1、dup_x2、dup2、dup2_x1和dup2_x2。初看这些指令，容易混淆而难以理解。经过分类和找规律，可以通过"指令系数法"来理解记忆，非常简单：
//不带_x的指令是复制栈顶数据并压入栈顶。包括两个指令，dup和dup2

//带_x的指令是复制栈顶数据并插入栈顶以下的某个位置。共有4个指令
//dup的系数代表要复制的Slot个数。
//dup开头的指令用于复制1个Slot的数据。例如1个int或1个reference类型数据
//dup2开头的指令用于复制2个Slot的数据。例如1个long，或2个int，或1个int+1个float类型数据
//
//对于带_x的复制插入指令，只要将指令的dup和x的系数相加，结果即为需要插入的位置。因此
//
//dup_x1插入位置：1+1=2，即栈顶2个Slot下面。
//dup_x2插入位置：1+2=3，即栈顶3个Slot下面。
//dup2_x1插入位置：2+1=3，即栈顶3个Slot下面。
//dup2_x2插入位置：2+2=4，即栈顶4个Slot下面。

//0x59 dup      复制栈顶数值并将复制值压入栈顶
//0x5a dup_x1    复制栈顶数值并将两个复制值压入栈顶
//0x5b dup_x2    复制栈顶数值并将三个（或两个）复制值压入栈顶

type Dup struct {
	base.NoOperandsInstruction
}

func (d *Dup) Execute(frame *runtimedata.Frame) {
	slot := frame.OperandStack().PopSlot()
	frame.OperandStack().PushSlot(slot)
	frame.OperandStack().PushSlot(slot)
}

type DupX1 struct {
	base.NoOperandsInstruction
}

func (d *DupX1) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

type DupX2 struct {
	base.NoOperandsInstruction
}

func (d *DupX2) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

//0x5c dup2     复制栈顶一个（long或double类型的)或两个（其它）数值并将复制值压入栈顶
//0x5d dup2_x1    dup_x1的双倍版本
//0x5e dup2_x2    dup_x2 的双倍版本
type Dup2 struct {
	base.NoOperandsInstruction
}

func (d *Dup2) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

type Dup2X1 struct {
	base.NoOperandsInstruction
}

func (d *Dup2X1) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

type Dup2X3 struct {
	base.NoOperandsInstruction
}

func (d *Dup2X3) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	slot4 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot4)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

package stack

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

// 0x5f swap     将栈最顶端的两个数值互换(数值不能是long或double类型的)
type Swap struct {
	base.NoOperandsInstruction
}

func (s *Swap) Execute(frame *runtimedata.Frame) {
	slot1 := frame.OperandStack().PopSlot()
	slot2 := frame.OperandStack().PopSlot()
	frame.OperandStack().PushSlot(slot1)
	frame.OperandStack().PushSlot(slot2)
}

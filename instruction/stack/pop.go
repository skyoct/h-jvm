package stack

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

// 0x57 pop      将栈顶数值弹出 (数值不能是long或double类型的)
type Pop struct {
	base.NoOperandsInstruction
}

func (p *Pop) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PopSlot()
}

// pop2     将栈顶的一个（long或double类型的)或两个数值弹出（其它）
type Pop2 struct {
	base.NoOperandsInstruction
}

func (p *Pop2) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PopSlot()
	frame.OperandStack().PopSlot()
}

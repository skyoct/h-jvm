package math

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

//0x82 ixor     将栈顶两int型数值作“按位异或”并将结果压入栈顶
//0x83 lxor     将栈顶两long型数值作“按位异或”并将结果压入栈顶

type IXor struct {
	base.NoOperandsInstruction
}

func (i *IXor) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopInt()
	val2 := stack.PopInt()
	stack.PushInt(val1 ^ val2)
}

type LXor struct {
	base.NoOperandsInstruction
}

func (l *LXor) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopLong()
	val2 := stack.PopLong()
	stack.PushLong(val1 ^ val2)
}

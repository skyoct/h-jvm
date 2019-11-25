package math

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

//0x80 ior      将栈顶两int型数值作“按位或”并将结果压入栈顶
//0x81 lor      将栈顶两long型数值作“按位或”并将结果压入栈顶

type IOr struct {
	base.NoOperandsInstruction
}

func (i *IOr) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopInt()
	val2 := stack.PopInt()
	stack.PushInt(val1 | val2)
}

type LOr struct {
	base.NoOperandsInstruction
}

func (l *LOr) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopLong()
	val2 := stack.PopLong()
	stack.PushLong(val1 | val2)
}

package math

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

//0x64 isub     将栈顶两int型数值相减并将结果压入栈顶
//0x65 lsub     将栈顶两long型数值相减并将结果压入栈顶
//0x66 fsub     将栈顶两float型数值相减并将结果压入栈顶
//0x67 dsub     将栈顶两double型数值相减并将结果压入栈顶

type ISub struct {
	base.NoOperandsInstruction
}

func (i *ISub) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopInt()
	val2 := stack.PopInt()
	stack.PushInt(val2 - val1)
}

type LSub struct {
	base.NoOperandsInstruction
}

func (l *LSub) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopLong()
	val2 := stack.PopLong()
	stack.PushLong(val2 - val1)
}

type FSub struct {
	base.NoOperandsInstruction
}

func (f *FSub) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopFloat()
	val2 := stack.PopFloat()
	stack.PushFloat(val2 - val1)
}

type DSub struct {
	base.NoOperandsInstruction
}

func (d *DSub) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopDouble()
	val2 := stack.PopDouble()
	stack.PushDouble(val2 - val1)
}

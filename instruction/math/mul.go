package math

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

//0x68 imul     将栈顶两int型数值相乘并将结果压入栈顶
//0x69 lmul     将栈顶两long型数值相乘并将结果压入栈顶
//0x6a fmul     将栈顶两float型数值相乘并将结果压入栈顶
//0x6b dmul     将栈顶两double型数值相乘并将结果压入栈顶

type IMul struct {
	base.NoOperandsInstruction
}

func (i *IMul) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopInt()
	val2 := stack.PopInt()
	stack.PushInt(val1 * val2)
}

type LMul struct {
	base.NoOperandsInstruction
}

func (l *LMul) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopLong()
	val2 := stack.PopLong()
	stack.PushLong(val1 * val2)
}

type FMul struct {
	base.NoOperandsInstruction
}

func (f *FMul) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopFloat()
	val2 := stack.PopFloat()
	stack.PushFloat(val1 * val2)
}

type DMul struct {
	base.NoOperandsInstruction
}

func (d *DMul) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopDouble()
	val2 := stack.PopDouble()
	stack.PushDouble(val1 * val2)
}

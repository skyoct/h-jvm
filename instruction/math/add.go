package math

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

//0x60 iadd     将栈顶两int型数值相加并将结果压入栈顶
//0x61 ladd     将栈顶两long型数值相加并将结果压入栈顶
//0x62 fadd     将栈顶两float型数值相加并将结果压入栈顶
//0x63 dadd     将栈顶两double型数值相加并将结果压入栈顶

type IAdd struct {
	base.NoOperandsInstruction
}

func (i *IAdd) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopInt()
	val2 := stack.PopInt()
	stack.PushInt(val1 + val2)
}

type LAdd struct {
	base.NoOperandsInstruction
}

func (l *LAdd) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopLong()
	val2 := stack.PopLong()
	stack.PushLong(val1 + val2)
}

type FAdd struct {
	base.NoOperandsInstruction
}

func (f *FAdd) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopFloat()
	val2 := stack.PopFloat()
	stack.PushFloat(val1 + val2)
}

type DAdd struct {
	base.NoOperandsInstruction
}

func (d *DAdd) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopDouble()
	val2 := stack.PopDouble()
	stack.PushDouble(val1 + val2)
}

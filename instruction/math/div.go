package math

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

//0x6c idiv     将栈顶两int型数值相除并将结果压入栈顶
//0x6d ldiv     将栈顶两long型数值相除并将结果压入栈顶
//0x6e fdiv     将栈顶两float型数值相除并将结果压入栈顶
//0x6f ddiv     将栈顶两double型数值相除并将结果压入栈顶

type IDiv struct {
	base.NoOperandsInstruction
}

func (i *IDiv) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopInt()
	val2 := stack.PopInt()
	stack.PushInt(val1 / val2)
}

type LDiv struct {
	base.NoOperandsInstruction
}

func (l *LDiv) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopLong()
	val2 := stack.PopLong()
	stack.PushLong(val1 / val2)
}

type FDiv struct {
	base.NoOperandsInstruction
}

func (f *FDiv) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopFloat()
	val2 := stack.PopFloat()
	stack.PushFloat(val1 / val2)
}

type DDiv struct {
	base.NoOperandsInstruction
}

func (d *DDiv) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopDouble()
	val2 := stack.PopDouble()
	stack.PushDouble(val1 / val2)
}

package math

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

//0x74 ineg     将栈顶int型数值取负并将结果压入栈顶
//0x75 lneg     将栈顶long型数值取负并将结果压入栈顶
//0x76 fneg     将栈顶float型数值取负并将结果压入栈顶
//0x77 dneg     将栈顶double型数值取负并将结果压入栈顶

type INeg struct {
	base.NoOperandsInstruction
}

func (i *INeg) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(-val)
}

type LNeg struct {
	base.NoOperandsInstruction
}

func (l *LNeg) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}

type FNeg struct {
	base.NoOperandsInstruction
}

func (f *FNeg) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}

type DNeg struct {
	base.NoOperandsInstruction
}

func (d *DNeg) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}

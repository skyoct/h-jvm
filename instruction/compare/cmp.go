package compare

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

//0x94 lcmp     比较栈顶两long型数值大小，并将结果（1，0，-1）压入栈顶
//0x95 fcmpl     比较栈顶两float型数值大小，并将结果（1，0，-1）压入栈顶；当其中一个数值为NaN时，将-1压入栈顶
//0x96 fcmpg     比较栈顶两float型数值大小，并将结果（1，0，-1）压入栈顶；当其中一个数值为NaN时，将1压入栈顶
//0x97 dcmpl     比较栈顶两double型数值大小，并将结果（1，0，-1）压入栈顶；当其中一个数值为NaN时，将-1压入栈顶
//0x98 dcmpg     比较栈顶两double型数值大小，并将结果（1，0，-1）压入栈顶；当其中一个数值为NaN时，将1压入栈顶
//
// 栈顶的小于栈顶下面的为1 相等的为0 否则为-1

type LCmp struct {
	base.NoOperandsInstruction
}

func (l *LCmp) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopLong()
	val2 := stack.PopLong()
	if val1 < val2 {
		stack.PushInt(1)
	} else if val1 == val2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}

func _fcmp(frame *runtimedata.Frame, flag bool) {
	stack := frame.OperandStack()
	val1 := stack.PopFloat()
	val2 := stack.PopFloat()
	if val1 < val2 {
		stack.PushInt(1)
	} else if val1 == val2 {
		stack.PushInt(0)
	} else if flag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

type FCmpL struct {
	base.NoOperandsInstruction
}

func (f *FCmpL) Execute(frame *runtimedata.Frame) {
	_fcmp(frame, false)
}

type FCmpG struct {
	base.NoOperandsInstruction
}

func (f *FCmpG) Execute(frame *runtimedata.Frame) {
	_fcmp(frame, true)
}

func _dcmp(frame *runtimedata.Frame, flag bool) {
	stack := frame.OperandStack()
	val1 := stack.PopDouble()
	val2 := stack.PopDouble()
	if val1 < val2 {
		stack.PushInt(1)
	} else if val1 == val2 {
		stack.PushInt(0)
	} else if flag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

type DCmpL struct {
	base.NoOperandsInstruction
}

func (f *DCmpL) Execute(frame *runtimedata.Frame) {
	_dcmp(frame, false)
}

type DCmpG struct {
	base.NoOperandsInstruction
}

func (f *DCmpG) Execute(frame *runtimedata.Frame) {
	_dcmp(frame, true)
}

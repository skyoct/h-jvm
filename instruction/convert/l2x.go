package convert

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

//0x88 l2i      将栈顶long型数值强制转换成int型数值并将结果压入栈顶
//0x89 l2f      将栈顶long型数值强制转换成float型数值并将结果压入栈顶
//0x8a l2d      将栈顶long型数值强制转换成double型数值并将结果压入栈顶

type L2I struct {
	base.NoOperandsInstruction
}

func (l *L2I) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	val2 := int32(val)
	stack.PushInt(val2)
}

type L2F struct {
	base.NoOperandsInstruction
}

func (l *L2F) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	val2 := float32(val)
	stack.PushFloat(val2)
}

type L2D struct {
	base.NoOperandsInstruction
}

func (l *L2D) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	val2 := float64(val)
	stack.PushDouble(val2)
}

package convert

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

//0x8b f2i      将栈顶float型数值强制转换成int型数值并将结果压入栈顶
//0x8c f2l      将栈顶float型数值强制转换成long型数值并将结果压入栈顶
//0x8d f2d      将栈顶float型数值强制转换成double型数值并将结果压入栈顶

type F2I struct {
	base.NoOperandsInstruction
}

func (f *F2I) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	val2 := int32(val)
	stack.PushInt(val2)
}

type F2L struct {
	base.NoOperandsInstruction
}

func (f *F2L) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	val2 := int64(val)
	stack.PushLong(val2)
}

type F2D struct {
	base.NoOperandsInstruction
}

func (f *F2D) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	val2 := float64(val)
	stack.PushDouble(val2)
}

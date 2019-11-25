package convert

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

//0x8e d2i      将栈顶double型数值强制转换成int型数值并将结果压入栈顶
//0x8f d2l      将栈顶double型数值强制转换成long型数值并将结果压入栈顶
//0x90 d2f      将栈顶double型数值强制转换成float型数值并将结果压入栈顶

type D2I struct {
	base.NoOperandsInstruction
}

func (d *D2I) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	val2 := int32(val)
	stack.PushInt(val2)
}

type D2F struct {
	base.NoOperandsInstruction
}

func (d *D2F) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	val2 := float32(val)
	stack.PushFloat(val2)
}

type D2L struct {
	base.NoOperandsInstruction
}

func (d *D2L) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	val2 := int64(val)
	stack.PushLong(val2)
}

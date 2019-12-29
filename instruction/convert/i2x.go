package convert

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

//0x85 i2l      将栈顶int型数值强制转换成long型数值并将结果压入栈顶
//0x86 i2f      将栈顶int型数值强制转换成float型数值并将结果压入栈顶
//0x87 i2d      将栈顶int型数值强制转换成double型数值并将结果压入栈顶

//0x91 i2b      将栈顶int型数值强制转换成byte型数值并将结果压入栈顶
//0x92 i2c      将栈顶int型数值强制转换成char型数值并将结果压入栈顶
//0x93 i2s      将栈顶int型数值强制转换成short型数值并将结果压入栈顶

type I2L struct {
	base.NoOperandsInstruction
}

func (i *I2L) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	val2 := int64(val)
	stack.PushLong(val2)
}

type I2F struct {
	base.NoOperandsInstruction
}

func (i *I2F) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	val2 := float32(val)
	stack.PushFloat(val2)
}

type I2D struct {
	base.NoOperandsInstruction
}

func (i *I2D) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	val2 := float64(val)
	stack.PushDouble(val2)
}

type I2C struct {
	base.NoOperandsInstruction
}

func (i *I2C) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	val2 := int32(uint16(val))
	stack.PushInt(val2)
}

type I2S struct {
	base.NoOperandsInstruction
}

func (i *I2S) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	val2 := int32(int16(val))
	stack.PushInt(val2)
}

type I2B struct {
	base.NoOperandsInstruction
}

func (i *I2B) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	val2 := int32(int8(val))
	stack.PushInt(val2)
}

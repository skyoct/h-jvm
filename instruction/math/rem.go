package math

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
	"math"
)

//0x70 irem     将栈顶两int型数值作取模运算并将结果压入栈顶
//0x71 lrem     将栈顶两long型数值作取模运算并将结果压入栈顶
//0x72 frem     将栈顶两float型数值作取模运算并将结果压入栈顶
//0x73 drem     将栈顶两double型数值作取模运算并将结果压入栈顶

type IRem struct {
	base.NoOperandsInstruction
}

func (i *IRem) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopInt()
	val2 := stack.PopInt()
	if val2 == 0 { // 被除数不能为0
		panic("java.lang.ArithmeticException: / by zero")
	}
	stack.PushInt(val1 % val2)
}

type LRem struct {
	base.NoOperandsInstruction
}

func (l *LRem) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopLong()
	val2 := stack.PopLong()
	if val2 == 0 { // 被除数不能为0
		panic("java.lang.ArithmeticException: / by zero")
	}
	stack.PushLong(val1 % val2)
}

type FRem struct {
	base.NoOperandsInstruction
}

func (f *FRem) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopFloat()
	val2 := stack.PopFloat()
	result := float32(math.Mod(float64(val1), float64(val2)))
	stack.PushFloat(result)
}

type DRem struct {
	base.NoOperandsInstruction
}

func (d *DRem) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopDouble()
	val2 := stack.PopDouble()
	result := math.Mod(val1, val2)
	stack.PushDouble(result)
}

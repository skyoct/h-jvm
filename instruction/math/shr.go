package math

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

//0x7a ishr     将int型数值右（符号）移位指定位数并将结果压入栈顶
//0x7b lshr     将long型数值右（符号）移位指定位数并将结果压入栈顶
//0x7c iushr     将int型数值右（无符号）移位指定位数并将结果压入栈顶
//0x7d lushr     将long型数值右（无符号）移位指定位数并将结果压入栈顶

type IShr struct {
	base.NoOperandsInstruction
}

func (i *IShr) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopInt() // 移动位数
	val2 := stack.PopInt()
	val := uint32(val1) & 0x1f
	result := val2 >> val
	stack.PushInt(result)
}

type LShr struct {
	base.NoOperandsInstruction
}

func (l *LShr) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopInt()
	val2 := stack.PopLong()
	val := uint32(val1) & 0x3f
	result := val2 >> val
	stack.PushLong(result)
}

type IUShr struct {
	base.NoOperandsInstruction
}

func (i *IUShr) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopInt() // 移动位数
	val2 := stack.PopInt()
	val := uint32(val1) & 0x1f
	result := int32(uint32(val2) >> val)
	stack.PushInt(result)
}

type LUShr struct {
	base.NoOperandsInstruction
}

func (l *LUShr) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopInt()
	val2 := stack.PopLong()
	val := uint32(val1) & 0x3f
	result := int64(uint64(val2) >> val)
	stack.PushLong(result)
}

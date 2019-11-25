package math

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

//0x78 ishl     将int型数值左移位指定位数并将结果压入栈顶
//0x79 lshl     将long型数值左移位指定位数并将结果压入栈顶

type IShl struct {
	base.NoOperandsInstruction
}

func (i *IShl) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopInt() // 移动位数
	val2 := stack.PopInt() // 需要移动的值
	val := uint32(val1) & 0x1f
	result := val2 << val
	stack.PushInt(result)
}

type LShl struct {
	base.NoOperandsInstruction
}

func (l *LShl) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val1 := stack.PopInt() // 移动位数
	val2 := stack.PopLong()
	val := uint32(val1) & 0x3f
	result := val2 << val
	stack.PushLong(result)
}

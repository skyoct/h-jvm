package reference

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

// 0xbe
// arraylength	获得数组的长度值并压入栈顶
// 只需要一个操作数，栈顶弹出的数组引用
// 计算完长度后 压入栈中

type ArrayLength struct {
	base.NoOperandsInstruction
}

func (a *ArrayLength) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	array := stack.PopRef()
	if array == nil{
		panic("java.lang.NullPointException")
	}
	length := array.ArrayLength()
	stack.PushInt(length) //长度压入栈中
}


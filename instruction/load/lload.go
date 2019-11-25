package load

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

func _lLoad(frame *runtimedata.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}

// 0x16 lload     将指定的long型本地变量推送至栈顶
type LLoad struct {
	base.Index8Instruction
}

func (l *LLoad) Execute(frame *runtimedata.Frame) {
	_lLoad(frame, l.Index)
}

//0x1e lload_0    将第一个long型本地变量推送至栈顶
//0x1f lload_1    将第二个long型本地变量推送至栈顶
//0x20 lload_2    将第三个long型本地变量推送至栈顶
//0x21 lload_3    将第四个long型本地变量推送至栈顶

type LLoad0 struct {
	base.NoOperandsInstruction
}

func (l *LLoad0) Execute(frame *runtimedata.Frame) {
	_lLoad(frame, 0)
}

type LLoad1 struct {
	base.NoOperandsInstruction
}

func (l *LLoad1) Execute(frame *runtimedata.Frame) {
	_lLoad(frame, 1)
}

type LLoad2 struct {
	base.NoOperandsInstruction
}

func (l *LLoad2) Execute(frame *runtimedata.Frame) {
	_lLoad(frame, 2)
}

type LLoad3 struct {
	base.NoOperandsInstruction
}

func (l *LLoad3) Execute(frame *runtimedata.Frame) {
	_lLoad(frame, 3)
}

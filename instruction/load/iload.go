package load

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

func _iLoad(frame *runtimedata.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

// 0x15 iload     将指定的int型本地变量推送至栈顶
type ILoad struct {
	base.Index8Instruction
}

func (i *ILoad) Execute(frame *runtimedata.Frame) {
	_iLoad(frame, i.Index)
}

// 0x1a iload_0    将第一个int型本地变量推送至栈顶
type ILoad0 struct {
	base.NoOperandsInstruction
}

func (i *ILoad0) Execute(frame *runtimedata.Frame) {
	_iLoad(frame, 0)
}

// 0x1a iload_1    将第二个int型本地变量推送至栈顶
type ILoad1 struct {
	base.NoOperandsInstruction
}

func (i *ILoad1) Execute(frame *runtimedata.Frame) {
	_iLoad(frame, 1)
}

type ILoad2 struct {
	base.NoOperandsInstruction
}

func (i *ILoad2) Execute(frame *runtimedata.Frame) {
	_iLoad(frame, 2)
}

type ILoad3 struct {
	base.NoOperandsInstruction
}

func (i *ILoad3) Execute(frame *runtimedata.Frame) {
	_iLoad(frame, 3)
}

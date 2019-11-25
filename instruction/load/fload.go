package load

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

func _fLoad(frame *runtimedata.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}

// 0x17 fload     将指定的float型本地变量推送至栈顶
type FLoad struct {
	base.Index8Instruction
}

func (f *FLoad) Execute(frame *runtimedata.Frame) {
	_fLoad(frame, f.Index)
}

//0x22 fload_0    将第一个float型本地变量推送至栈顶
//0x23 fload_1    将第二个float型本地变量推送至栈顶
//0x24 fload_2    将第三个float型本地变量推送至栈顶
//0x25 fload_3    将第四个float型本地变量推送至栈顶

type FLoad0 struct {
	base.NoOperandsInstruction
}

func (f *FLoad0) Execute(frame *runtimedata.Frame) {
	_fLoad(frame, 0)
}

type FLoad1 struct {
	base.NoOperandsInstruction
}

func (f *FLoad1) Execute(frame *runtimedata.Frame) {
	_fLoad(frame, 1)
}

type FLoad2 struct {
	base.NoOperandsInstruction
}

func (f *FLoad2) Execute(frame *runtimedata.Frame) {
	_fLoad(frame, 2)
}

type FLoad3 struct {
	base.NoOperandsInstruction
}

func (f *FLoad3) Execute(frame *runtimedata.Frame) {
	_fLoad(frame, 3)
}

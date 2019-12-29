package load

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

func _dLoad(frame *runtimedata.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}

// 0x18 dload     将指定的double型本地变量推送至栈顶
type DLoad struct {
	base.Index8Instruction
}

func (d *DLoad) Execute(frame *runtimedata.Frame) {
	_dLoad(frame, d.Index)
}

// 0x26 dload_0    将第一个double型本地变量推送至栈顶
//0x27 dload_1    将第二个double型本地变量推送至栈顶
//0x28 dload_2    将第三个double型本地变量推送至栈顶
//0x29 dload_3    将第四个double型本地变量推送至栈顶

type DLoad0 struct {
	base.NoOperandsInstruction
}

func (d *DLoad0) Execute(frame *runtimedata.Frame) {
	_dLoad(frame, 0)
}

type DLoad1 struct {
	base.NoOperandsInstruction
}

func (d *DLoad1) Execute(frame *runtimedata.Frame) {
	_dLoad(frame, 1)
}

type DLoad2 struct {
	base.NoOperandsInstruction
}

func (d *DLoad2) Execute(frame *runtimedata.Frame) {
	_dLoad(frame, 2)
}

type DLoad3 struct {
	base.NoOperandsInstruction
}

func (d *DLoad3) Execute(frame *runtimedata.Frame) {
	_dLoad(frame, 3)
}

package load

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

func _aLoad(frame *runtimedata.Frame, index uint) {
	ref := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(ref)
}

// 0x19 aload     将指定的引用类型本地变量推送至栈顶
type ALoad struct {
	base.Index8Instruction
}

func (a *ALoad) Execute(frame *runtimedata.Frame) {
	_aLoad(frame, a.Index)
}

// 0x2a aload_0    将第一个引用类型本地变量推送至栈顶
//0x2b aload_1    将第二个引用类型本地变量推送至栈顶
//0x2c aload_2    将第三个引用类型本地变量推送至栈顶
//0x2d aload_3    将第四个引用类型本地变量推送至栈顶

type ALoad0 struct {
	base.NoOperandsInstruction
}

func (a *ALoad0) Execute(frame *runtimedata.Frame) {
	_aLoad(frame, 0)
}

type ALoad1 struct {
	base.NoOperandsInstruction
}

func (a *ALoad1) Execute(frame *runtimedata.Frame) {
	_aLoad(frame, 1)
}

type ALoad2 struct {
	base.NoOperandsInstruction
}

func (a *ALoad2) Execute(frame *runtimedata.Frame) {
	_aLoad(frame, 2)
}

type ALoad3 struct {
	base.NoOperandsInstruction
}

func (a *ALoad3) Execute(frame *runtimedata.Frame) {
	_aLoad(frame, 3)
}

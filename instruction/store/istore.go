package store

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

func _iStore(frame *runtimedata.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}

//将栈顶int型数值存入指定本地变量
type IStore struct {
	base.Index8Instruction
}

func (i *IStore) Execute(frame *runtimedata.Frame) {
	_iStore(frame, i.Index)
}

//0x3b istore_0   将栈顶int型数值存入第一个本地变量
//0x3c istore_1   将栈顶int型数值存入第二个本地变量
//0x3d istore_2   将栈顶int型数值存入第三个本地变量
//0x3e istore_3   将栈顶int型数值存入第四个本地变量

type IStore0 struct {
	base.NoOperandsInstruction
}

func (i *IStore0) Execute(frame *runtimedata.Frame) {
	_iStore(frame, 0)
}

type IStore1 struct {
	base.NoOperandsInstruction
}

func (i *IStore1) Execute(frame *runtimedata.Frame) {
	_iStore(frame, 1)
}

type IStore2 struct {
	base.NoOperandsInstruction
}

func (i *IStore2) Execute(frame *runtimedata.Frame) {
	_iStore(frame, 2)
}

type IStore3 struct {
	base.NoOperandsInstruction
}

func (i *IStore3) Execute(frame *runtimedata.Frame) {
	_iStore(frame, 3)
}

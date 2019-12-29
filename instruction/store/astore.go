package store

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

func _aStore(frame *runtimedata.Frame, index uint) {
	ref := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, ref)
}

// 0x3a astore    将栈顶引用型数值存入指定本地变量
type AStore struct {
	base.Index8Instruction
}

func (a *AStore) Execute(frame *runtimedata.Frame) {
	_aStore(frame, a.Index)
}

//0x4b astore_0   将栈顶引用型数值存入第一个本地变量
//0x4c astore_1   将栈顶引用型数值存入第二个本地变量
//0x4d astore_2   将栈顶引用型数值存入第三个本地变量
//0x4e astore_3   将栈顶引用型数值存入第四个本地变量

type AStore0 struct {
	base.NoOperandsInstruction
}

func (a *AStore0) Execute(frame *runtimedata.Frame) {
	_aStore(frame, 0)
}

type AStore1 struct {
	base.NoOperandsInstruction
}

func (a *AStore1) Execute(frame *runtimedata.Frame) {
	_aStore(frame, 1)
}

type AStore2 struct {
	base.NoOperandsInstruction
}

func (a *AStore2) Execute(frame *runtimedata.Frame) {
	_aStore(frame, 2)
}

type AStore3 struct {
	base.NoOperandsInstruction
}

func (a *AStore3) Execute(frame *runtimedata.Frame) {
	_aStore(frame, 3)
}

package store

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

func _fStore(frame *runtimedata.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}

// 0x38 fstore    将栈顶float型数值存入指定本地变量
type FStore struct {
	base.Index8Instruction
}

func (f *FStore) Execute(frame *runtimedata.Frame) {
	_fStore(frame, f.Index)
}

//0x43 fstore_0   将栈顶float型数值存入第一个本地变量
//0x44 fstore_1   将栈顶float型数值存入第二个本地变量
//0x45 fstore_2   将栈顶float型数值存入第三个本地变量
//0x46 fstore_3   将栈顶float型数值存入第四个本地变量

type FStore0 struct {
	base.NoOperandsInstruction
}

func (f *FStore0) Execute(frame *runtimedata.Frame) {
	_fStore(frame, 0)
}

type FStore1 struct {
	base.NoOperandsInstruction
}

func (f *FStore1) Execute(frame *runtimedata.Frame) {
	_fStore(frame, 1)
}

type FStore2 struct {
	base.NoOperandsInstruction
}

func (f *FStore2) Execute(frame *runtimedata.Frame) {
	_fStore(frame, 2)
}

type FStore3 struct {
	base.NoOperandsInstruction
}

func (f *FStore3) Execute(frame *runtimedata.Frame) {
	_fStore(frame, 3)
}

package store

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

func _dStore(frame *runtimedata.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}

//0x39 dstore    将栈顶double型数值存入指定本地变量

type DStore struct {
	base.Index8Instruction
}

func (d *DStore) Execute(frame *runtimedata.Frame) {
	_dStore(frame, d.Index)
}

// 0x47 dstore_0   将栈顶double型数值存入第一个本地变量
//0x48 dstore_1   将栈顶double型数值存入第二个本地变量
//0x49 dstore_2   将栈顶double型数值存入第三个本地变量
//0x4a dstore_3   将栈顶double型数值存入第四个本地变量

type DStore0 struct {
	base.NoOperandsInstruction
}

func (d *DStore0) Execute(frame *runtimedata.Frame) {
	_dStore(frame, 0)
}

type DStore1 struct {
	base.NoOperandsInstruction
}

func (d *DStore1) Execute(frame *runtimedata.Frame) {
	_dStore(frame, 1)
}

type DStore2 struct {
	base.NoOperandsInstruction
}

func (d *DStore2) Execute(frame *runtimedata.Frame) {
	_dStore(frame, 2)
}

type DStore3 struct {
	base.NoOperandsInstruction
}

func (d *DStore3) Execute(frame *runtimedata.Frame) {
	_dStore(frame, 3)
}

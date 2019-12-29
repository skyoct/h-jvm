package store

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

func _lStore(frame *runtimedata.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

// 0x37 lstore    将栈顶long型数值存入指定本地变量
type LStore struct {
	base.Index8Instruction
}

func (l *LStore) Execute(frame *runtimedata.Frame) {
	_lStore(frame, l.Index)
}

//0x3f lstore_0   将栈顶long型数值存入第一个本地变量
//0x40 lstore_1   将栈顶long型数值存入第二个本地变量
//0x41 lstore_2   将栈顶long型数值存入第三个本地变量
//0x42 lstore_3   将栈顶long型数值存入第四个本地变量

type LStore0 struct {
	base.NoOperandsInstruction
}

func (l *LStore0) Execute(frame *runtimedata.Frame) {
	_lStore(frame, 0)
}

type LStore1 struct {
	base.NoOperandsInstruction
}

func (l *LStore1) Execute(frame *runtimedata.Frame) {
	_lStore(frame, 1)
}

type LStore2 struct {
	base.NoOperandsInstruction
}

func (l *LStore2) Execute(frame *runtimedata.Frame) {
	_lStore(frame, 2)
}

type LStore3 struct {
	base.NoOperandsInstruction
}

func (l *LStore3) Execute(frame *runtimedata.Frame) {
	_lStore(frame, 3)
}

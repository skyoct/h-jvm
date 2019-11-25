package base

import "h-jvm/runtimedata"

type Instruction interface {
	FetchOperands(reader *CodeReader)
	Execute(frame *runtimedata.Frame)
}

// 没有操作数的指令
type NoOperandsInstruction struct {
}

// FetchOperands 为空
func (n *NoOperandsInstruction) FetchOperands(reader *CodeReader) {
}

// 跳转类型的指令
type BranchInstruction struct {
	Offset int // 偏移量
}

func (b *BranchInstruction) FetchOperands(reader *CodeReader) {
	b.Offset = int(reader.ReadInt16())
}

// 从局部变量表中获取数值
type Index8Instruction struct {
	Index uint
}

func (i *Index8Instruction) FetchOperands(reader *CodeReader) {
	i.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint
}

func (i *Index16Instruction) FetchOperands(reader *CodeReader) {
	i.Index = uint(reader.ReadUint16())
}

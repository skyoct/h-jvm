package control

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

//0xa7 goto     无条件跳转
type Goto struct {
	base.BranchInstruction
}

func (g *Goto) Execute(frame *runtimedata.Frame) {
	base.Branch(frame, g.Offset)
}

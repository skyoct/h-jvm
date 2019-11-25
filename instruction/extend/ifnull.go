package extend

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

//0xc6 ifnull    为null时跳转
//0xc7 ifnonnull   不为null时跳转

type IFNull struct {
	base.BranchInstruction
}

func (i *IFNull) Execute(frame *runtimedata.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, i.Offset)
	}
}

type IFNonNull struct {
	base.BranchInstruction
}

func (i *IFNonNull) Execute(frame *runtimedata.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, i.Offset)
	}
}

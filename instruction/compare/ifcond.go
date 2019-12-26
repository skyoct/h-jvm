package compare

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

//0x99 ifeq     当栈顶int型数值等于0时跳转
//0x9a ifne     当栈顶int型数值不等于0时跳转
//0x9b iflt     当栈顶int型数值小于0时跳转
//0x9c ifge     当栈顶int型数值大于等于0时跳转
//0x9d ifgt     当栈顶int型数值大于0时跳转
//0x9e ifle     当栈顶int型数值小于等于0时跳转

type IFeq struct {
	base.BranchInstruction
}

func (i *IFeq) Execute(frame *runtimedata.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, i.Offset)
	}
}

type IFne struct {
	base.BranchInstruction
}

func (i *IFne) Execute(frame *runtimedata.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, i.Offset)
	}
}

type IFlt struct {
	base.BranchInstruction
}

func (i *IFlt) Execute(frame *runtimedata.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, i.Offset)
	}
}

type IFle struct {
	base.BranchInstruction
}

func (i *IFle) Execute(frame *runtimedata.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, i.Offset)
	}
}

type IFgt struct {
	base.BranchInstruction
}

func (i *IFgt) Execute(frame *runtimedata.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, i.Offset)
	}
}

type IFge struct {
	base.BranchInstruction
}

func (i *IFge) Execute(frame *runtimedata.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, i.Offset)
	}
}

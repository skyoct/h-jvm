package compare

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
	"h-jvm/runtimedata/metaspace"
)

func _iPop(frame *runtimedata.Frame) (int32, int32) {
	stack := frame.OperandStack()
	val1 := stack.PopInt()
	val2 := stack.PopInt()
	return val1, val2
}

func _aPop(frame *runtimedata.Frame) (*metaspace.Object, *metaspace.Object) {
	ref1 := frame.OperandStack().PopRef()
	ref2 := frame.OperandStack().PopRef()
	return ref1, ref2
}

//0x9f if_icmpeq   比较栈顶两int型数值大小，当结果等于0时跳转
//0xa0 if_icmpne   比较栈顶两int型数值大小，当结果不等于0时跳转
//0xa1 if_icmplt   比较栈顶两int型数值大小，当结果小于0时跳转
//0xa2 if_icmpge   比较栈顶两int型数值大小，当结果大于等于0时跳转
//0xa3 if_icmpgt   比较栈顶两int型数值大小，当结果大于0时跳转
//0xa4 if_icmple   比较栈顶两int型数值大小，当结果小于等于0时跳转
// 栈顶元素为v1 栈顶下面的元素为v2  v1 < v2 == 1  v1 > v2 == -1

type IFICmpEq struct {
	base.BranchInstruction
}

func (i *IFICmpEq) Execute(frame *runtimedata.Frame) {
	if val1, val2 := _iPop(frame); val1 == val2 {
		base.Branch(frame, i.Offset)
	}
}

type IFICmpNe struct {
	base.BranchInstruction
}

func (i *IFICmpNe) Execute(frame *runtimedata.Frame) {
	if val1, val2 := _iPop(frame); val1 != val2 {
		base.Branch(frame, i.Offset)
	}
}

type IFICmpLt struct {
	base.BranchInstruction
}

func (i *IFICmpLt) Execute(frame *runtimedata.Frame) {
	if val1, val2 := _iPop(frame); val1 > val2 {
		base.Branch(frame, i.Offset)
	}
}

type IFICmpLe struct {
	base.BranchInstruction
}

func (i *IFICmpLe) Execute(frame *runtimedata.Frame) {
	if val1, val2 := _iPop(frame); val1 >= val2 {
		base.Branch(frame, i.Offset)
	}
}

type IFICmpGt struct {
	base.BranchInstruction
}

func (i *IFICmpGt) Execute(frame *runtimedata.Frame) {
	if val1, val2 := _iPop(frame); val1 < val2 {
		base.Branch(frame, i.Offset)
	}
}

type IFICmpGe struct {
	base.BranchInstruction
}

func (i *IFICmpGe) Execute(frame *runtimedata.Frame) {
	if val1, val2 := _iPop(frame); val1 <= val2 {
		base.Branch(frame, i.Offset)
	}
}

//0xa5 if_acmpeq   比较栈顶两引用型数值，当结果相等时跳转
//0xa6 if_acmpne   比较栈顶两引用型数值，当结果不相等时跳转

type IFAcmEq struct {
	base.BranchInstruction
}

func (i *IFAcmEq) Execute(frame *runtimedata.Frame) {
	if ref1, ref2 := _aPop(frame); ref1 == ref2 {
		base.Branch(frame, i.Offset)
	}
}

type IFAcmNe struct {
	base.BranchInstruction
}

func (i *IFAcmNe) Execute(frame *runtimedata.Frame) {
	if ref1, ref2 := _aPop(frame); ref1 != ref2 {
		base.Branch(frame, i.Offset)
	}
}

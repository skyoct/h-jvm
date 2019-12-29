package reference

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
	"h-jvm/runtimedata/metaspace"
)

// 0xc0 checkcast 检验类型转换，检验未通过将抛出ClassCastException
//checkcast指令和instanceof指令很像，区别在于：instanceof指令会改变操作数栈（弹出对象引用，推入判断结果）；
// checkcast则不改 变操作数栈（如果判断失败，直接抛出ClassCastException异常）
type CheckCast struct {
	base.Index16Instruction
}

func (c *CheckCast) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)
	if ref == nil {
		return
	}
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(c.Index).(*metaspace.ClassRef)
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}

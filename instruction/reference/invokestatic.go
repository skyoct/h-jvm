package reference

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
	"h-jvm/runtimedata/metaspace"
)

type InvokeStatic struct {
	base.Index16Instruction
}

func (i * InvokeStatic) Execute(frame *runtimedata.Frame) {
	cp := frame.Method().Class().ConstantPool() // 获取常量池
	methodRef := cp.GetConstant(i.Index).(*metaspace.MethodRef) // 获取方法的符号引用
	resolveMethod := methodRef.ResolvedMethod()
	if !resolveMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	base.InvokeMethod(frame, resolveMethod)
}






package reference

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
	"h-jvm/runtimedata/metaspace"
)

// 0xb7 invokespecial   调用超类构造方法，实例初始化方法，私有方法

type InvokeSpecial struct {
	base.Index16Instruction
}

func (i *InvokeSpecial) Execute(frame *runtimedata.Frame) {
	cp := frame.Method().Class().ConstantPool() // 获取当前类的常量池
	methodRef := cp.GetConstant(i.Index).(*metaspace.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	base.InvokeMethod(frame, resolvedMethod)
}

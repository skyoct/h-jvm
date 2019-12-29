package reference

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
	"h-jvm/runtimedata/metaspace"
)

// instanceof指令判断对象是否是某个类的实例（或者对象的类是否实现了某个接口）
// 0xc1 instanceof 检验对象是否是指定的类的实例，如果是将1压入栈顶，否则将0压入栈顶
type InstanceOf struct {
	base.Index16Instruction
}

func (i *InstanceOf) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil { // 不是 直接把0压入栈顶
		stack.PushInt(0)
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(i.Index).(*metaspace.ClassRef)
	class := classRef.ResolvedClass()
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}

}

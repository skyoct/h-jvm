package reference

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
	"h-jvm/runtimedata/metaspace"
)

//new指令的操作数是一个uint16索引，来自字节码。通过这个索引，
// 可以从当前类的运行时常量池中找到一个类符号引用。
// 解析这个类符号引用，拿到类数据，然后创建对象，并把对象引用推入栈 顶，new指令的工作就完成了。
type New struct {
	base.Index16Instruction
}

func (n *New) Execute(frame *runtimedata.Frame) {
	cp := frame.Method().Class().ConstantPool() // 得到这个栈帧所属方法的类的
	classRef := cp.GetConstant(n.Index).(*metaspace.ClassRef)
	class := classRef.ResolvedClass()              // 加载类
	if !class.InitStarted() {
		frame.RevertNextPc() // 把pc指向当前指令（回退一步）
		base.InitClass(frame.Thread(), class)
		return
	}
	if class.IsInterface() || class.IsAbstract() { // 如果是接口或者抽象类直接报错
		panic("java.lang.InstantiationError")
	}
	ref := class.NewObject() // 创建一个对象
	frame.OperandStack().PushRef(ref)
}

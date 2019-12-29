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
	currentClass := frame.Method().Class()  //获取当前类
	cp := currentClass.ConstantPool() // 获取当前类的常量池
	methodRef := cp.GetConstant(i.Index).(*metaspace.MethodRef)
	resolvedClass := methodRef.ResolvedClass()
	resolvedMethod := methodRef.ResolvedMethod()
	if resolvedMethod.Name() == "<init>" && resolvedMethod.Class() != resolvedClass{
		panic("java.lang.NoSuchMethodError")
	}
	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 获取到要执行方法的实例对象的引用
	ref := frame.OperandStack().GetFromTop(resolvedMethod.ArgsSlotCount()-1)
	if ref == nil{
		panic("java.lang.NullPointException")
	}
	// 如果是protected 那么判断当前类是否是要执行类的子类
	if resolvedMethod.IsProtected() &&
		resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&
		!ref.Class().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")

	}
	methodToBeInvoke := metaspace.LookupMethodInClass(ref.Class(), resolvedMethod.Name(), resolvedMethod.Descriptor())
	if methodToBeInvoke == nil || methodToBeInvoke.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	base.InvokeMethod(frame, resolvedMethod)
}

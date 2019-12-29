package reference

import (
	"fmt"
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
	"h-jvm/runtimedata/metaspace"
)

// 0xb6 invokevirtual   调用实例方法

type InvokeVirtual struct {
	base.Index16Instruction
}

func (i *InvokeVirtual) Execute(frame *runtimedata.Frame) {
	currentClass := frame.Method().Class()
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(i.Index).(*metaspace.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 获取到要执行方法的实例对象的引用
	ref := frame.OperandStack().GetFromTop(resolvedMethod.ArgsSlotCount()-1)
	if ref == nil {
		// hack!
		if methodRef.Name() == "println" {
			_println(frame.OperandStack(), methodRef.Descriptor())
			return
		}

		panic("java.lang.NullPointerException")
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


// hack
func _println(stack *runtimedata.OperandStack, descriptor string)  {

	switch descriptor {
	case "(Z)V":
		fmt.Printf("%v\n", stack.PopInt() != 0)
	case "(C)V":
		fmt.Printf("%c\n", stack.PopInt())
	case "(I)V", "(B)V", "(S)V":
		fmt.Printf("%v\n", stack.PopInt())
	case "(F)V":
		fmt.Printf("%v\n", stack.PopFloat())
	case "(J)V":
		fmt.Printf("%v\n", stack.PopLong())
	case "(D)V":
		fmt.Printf("%v\n", stack.PopDouble())
	default:
		panic("println: " + descriptor)
	}
	stack.PopRef()
}


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
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(i.Index).(*metaspace.MethodRef)
	if methodRef.Name() == "println" {
		stack := frame.OperandStack()
		switch methodRef.Descriptor() {
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
			panic("println: " + methodRef.Descriptor())
		}
		stack.PopRef()
	}
}

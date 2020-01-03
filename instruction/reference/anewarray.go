package reference

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
	"h-jvm/runtimedata/metaspace"
)

// 0xbd
// anewarray创建一个引用型（如类，接口，数组）的数组，并将其引用值压入栈顶
// 两个操作数，第一个操作数来自字节码，通过这个可以在常量池找到类的符号引用
// 第二个操作数在栈顶

type ANewArray struct {
	base.Index16Instruction
}

func (a *ANewArray) Execute(frame *runtimedata.Frame) {
	cp := frame.Method().Class().ConstantPool()  // 获取常量池
	classRef := cp.GetConstant(a.Index).(metaspace.ClassRef) // 在常量池中查找符号引用
	class := classRef.ResolvedClass()  // 加载类
	count := frame.OperandStack().PopInt() // 获取栈顶的第一个值
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}
	arrayClass := class.ArrayClass()  //加载类
	array := arrayClass.NewArray(uint(count)) // 创建对象
	frame.OperandStack().PushRef(array) //  压入栈中
}



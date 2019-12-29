package reference

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
	"h-jvm/runtimedata/metaspace"
)

type InvokeInterface struct {
	index uint
}

// invokeinterface的操作操作码后面跟着四个字节，其中两个是常量池索引，1格式参数个数（我们通过计算得出）
// 还有一保存 值默认为0

func (i *InvokeInterface) FetchOperands(reader *base.CodeReader) {
	i.index = uint(reader.ReadInt16())
	reader.ReadInt8()
	reader.ReadInt8()
}

func (i *InvokeInterface) Execute(frame *runtimedata.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(i.index).(metaspace.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if resolvedMethod.IsStatic() || resolvedMethod.IsPrivate(){
		panic("java.lang.IncompatibleClassChangeError")
	}
	ref := frame.OperandStack().GetFromTop(resolvedMethod.ArgsSlotCount()-1)
	if ref == nil{
		panic("java.lang.NullPointException")
	}
	if !ref.Class().IsImplements(methodRef.ResolvedClass()) {
		panic("java.lang.IncompatibleClassChangeError")
	}
	methodToBeInvoke := metaspace.LookupMethodInClass(ref.Class(), resolvedMethod.Name(), resolvedMethod.Descriptor())
	if methodToBeInvoke == nil || methodToBeInvoke.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	if !methodToBeInvoke.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}
	base.InvokeMethod(frame, methodToBeInvoke)

}



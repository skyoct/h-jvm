package reference

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
	"h-jvm/runtimedata/metaspace"
)

// putstatic 的主要作用是给某个静态变量赋值
type PutStatic struct {
	base.Index16Instruction
}

func (p *PutStatic) Execute(frame *runtimedata.Frame) {
	currentMethod := frame.Method()
	currentClass := frame.Method().Class() // 当前类
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConstant(p.Index).(*metaspace.FieldRef) // 字段引用
	field := fieldRef.ResolvedField()                         // 解决类加载

	class := field.Class()
	if !class.InitStarted() {
		frame.RevertNextPc() // 把pc指向当前指令（回退一步）
		base.InitClass(frame.Thread(), class)
		return
	}
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError") // 如果不是静态的字段 抛出 uncompatibleClassChangeError
	}
	if field.IsFinal() { // 带final的类变量是无法修改的 其初始化阶段在类加载阶段连接的准备阶段初始化的
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}
	descriptor := field.Descriptor()
	slotId := field.SlotId() // 得到在槽中的位置
	slots := class.StaticVars()
	stack := frame.OperandStack()
	switch descriptor[0] { // 把操作数栈内的数据放到类的变量槽中（类变量）
	case 'Z', 'B', 'C', 'S', 'I':
		slots.SetInt(slotId, stack.PopInt())
	case 'F':
		slots.SetFloat(slotId, stack.PopFloat())
	case 'J':
		slots.SetLong(slotId, stack.PopLong())
	case 'D':
		slots.SetDouble(slotId, stack.PopDouble())
	case 'L', '[':
		slots.SetRef(slotId, stack.PopRef())
	}

}

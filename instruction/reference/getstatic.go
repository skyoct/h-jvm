package reference

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
	"h-jvm/runtimedata/metaspace"
)

type GetStatic struct {
	base.Index16Instruction
}

func (g *GetStatic) Execute(frame *runtimedata.Frame) {
	cp := frame.Method().Class().ConstantPool()
	fieldRef := cp.GetConstant(g.Index).(*metaspace.FieldRef)
	field := fieldRef.ResolvedField() // 通过字段引用加载字段
	class := field.Class()            // 得到字段所属的class
	
	if !class.InitStarted() {
		frame.RevertNextPc() // 把pc指向当前指令（回退一步）
		base.InitClass(frame.Thread(), class)
		return
	}
	if !field.IsStatic() {            // 如果当前字段不是静态的
		panic("java.lang.IncompatibleClassChangeError")
	}
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars() // 得到class类变量槽的引用
	stack := frame.OperandStack()
	switch descriptor[0] { // 从class类的变量槽中取除变量然后压入操作数栈中
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))
	}
}

package reference

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
	"h-jvm/runtimedata/metaspace"
)

type GetField struct {
	base.Index16Instruction
}

func (g *GetField) Execute(frame *runtimedata.Frame) {
	currentMethod := frame.Method()                           //获取当前帧所属的方法
	currentClass := currentMethod.Class()                     // 当前方法所属的类
	cp := currentClass.ConstantPool()                         // 获取常量池
	fieldRef := cp.GetConstant(g.Index).(*metaspace.FieldRef) // 获取字段引用
	field := fieldRef.ResolvedField()
	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := ref.Fields() // 获取对象实例的实例字段保存的
	switch descriptor[0] {
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
	default:
	}

}

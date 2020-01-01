package constant

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

//0x12 ldc      将int, float或String型常量值从常量池中推送至栈顶
//0x13 ldc_w     将int, float或String型常量值从常量池中推送至栈顶（宽索引）
//0x14 ldc2_w    将long或double型常量值从常量池中推送至栈顶（宽索引）

type Ldc struct {
	base.Index8Instruction
}

func (l *Ldc) Execute(frame *runtimedata.Frame) {
	_ldc(frame, l.Index)
}

type LdcW struct {
	base.Index8Instruction
}

func (l *LdcW) Execute(frame *runtimedata.Frame) {
	_ldc(frame, l.Index)
}

type Ldc2W struct {
	base.Index8Instruction
}

func (l *Ldc2W) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(l.Index)
	switch c.(type) {
	case int64:
		stack.PushLong(c.(int64))
	case float64:
		stack.PushDouble(c.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}

func _ldc(frame *runtimedata.Frame, index uint) {
	stack := frame.OperandStack()
	cp := frame.Method().Class().ConstantPool()
	c := cp.GetConstant(index)

	switch c.(type) {
	case int32:
		stack.PushInt(c.(int32))
	case float32:
		stack.PushFloat(c.(float32))
	// case string:
	// case *heap.ClassRef:
	// case MethodType, MethodHandle
	default:
		panic("todo: ldc!")
	}
}

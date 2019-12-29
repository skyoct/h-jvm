package control

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

//0xac ireturn    从当前方法返回int
//0xad lreturn    从当前方法返回long
//0xae freturn    从当前方法返回float
//0xaf dreturn    从当前方法返回double
//0xb0 areturn    从当前方法返回对象引用
//0xb1 return    从当前方法返回void

type Return struct {
	base.NoOperandsInstruction
}

// return 指令返回void  只需要将当前栈帧弹出
func (r *Return) Execute(frame *runtimedata.Frame) {
	frame.Thread().PopFrame()
}

// ireturn指令会将当前栈帧的操作数栈的栈顶元素弹出 然后将当前栈弹出，在把之前弹出的元素压入栈顶
type IReturn struct {
	base.NoOperandsInstruction
}

func (i *IReturn) Execute(frame *runtimedata.Frame) {
	thread := frame.Thread()
	val := frame.OperandStack().PopInt()
	// 弹出之前的栈帧
	thread.PopFrame()
	invokerFrame := thread.CurrentFrame() // 获取当前栈顶的栈帧
	invokerFrame.OperandStack().PushInt(val)
}

type LReturn struct {
	base.NoOperandsInstruction
}

func (l *LReturn) Execute(frame *runtimedata.Frame) {
	thread := frame.Thread()
	val := frame.OperandStack().PopLong()
	// 弹出之前的栈帧
	thread.PopFrame()
	invokerFrame := thread.CurrentFrame() // 获取当前栈顶的栈帧
	invokerFrame.OperandStack().PushLong(val)
}

type FReturn struct {
	base.NoOperandsInstruction
}

func (f *FReturn) Execute(frame *runtimedata.Frame) {
	thread := frame.Thread()
	val := frame.OperandStack().PopFloat()
	// 弹出之前的栈帧
	thread.PopFrame()
	invokerFrame := thread.CurrentFrame() // 获取当前栈顶的栈帧
	invokerFrame.OperandStack().PushFloat(val)
}

type DReturn struct {
	base.NoOperandsInstruction
}

func (d *DReturn) Execute(frame *runtimedata.Frame) {
	thread := frame.Thread()
	val := frame.OperandStack().PopDouble()
	// 弹出之前的栈帧
	thread.PopFrame()
	invokerFrame := thread.CurrentFrame() // 获取当前栈顶的栈帧
	invokerFrame.OperandStack().PushDouble(val)
}

type AReturn struct {
	base.NoOperandsInstruction
}

func (a *AReturn) Execute(frame *runtimedata.Frame) {
	thread := frame.Thread()
	val := frame.OperandStack().PopRef()
	// 弹出之前的栈帧
	thread.PopFrame()
	invokerFrame := thread.CurrentFrame() // 获取当前栈顶的栈帧
	invokerFrame.OperandStack().PushRef(val)
}

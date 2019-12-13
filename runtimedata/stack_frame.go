package runtimedata

import "h-jvm/runtimedata/metaspace"

// 栈帧 每个栈帧包括局部变量表和操作数栈
type Frame struct {
	next         *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	nextPc       int
	method       *metaspace.Method // 指向方法的指针 每个帧都是一个方法
}

func NewFrame(maxLocals, maxStack uint, method *metaspace.Method) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
		method:       method,
	}
}

func (f *Frame) OperandStack() *OperandStack {
	return f.operandStack
}

func (f *Frame) LocalVars() LocalVars {
	return f.localVars
}

func (f *Frame) NextPc() int {
	return f.nextPc
}

func (f *Frame) SetNextPc(nextPc int) {
	f.nextPc = nextPc
}

func (f *Frame) Thread() *Thread {
	return f.thread
}

func (f *Frame) Method() *metaspace.Method {
	return f.method
}

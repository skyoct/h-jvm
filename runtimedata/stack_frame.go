package runtimedata

// 栈帧 每个栈帧包括局部变量表和操作数栈
type Frame struct {
	next         *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	nextPc       int
}

func (f *Frame) OperandStack() *OperandStack {
	return f.operandStack
}

func (f *Frame) LocalVars() LocalVars {
	return f.localVars
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
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

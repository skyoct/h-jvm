package runtimedata

import "h-jvm/runtimedata/metaspace"

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024), // 虚拟机栈的最大深度
	}
}

// 放入frame
func (t *Thread) PushFrame(frame *Frame) {
	frame.thread = t // 把栈帧放入栈的时候 栈帧设置指向栈的值
	t.stack.push(frame)
}

// 当前frame
func (t *Thread) CurrentFrame() *Frame {
	return t.stack.top()
}

// 栈顶栈帧出栈
func (t *Thread) PopFrame() *Frame {
	return t.stack.pop()
}

func (t *Thread) SetPc(nextPc int) {
	t.pc = nextPc
}

func (t *Thread) Pc() int {
	return t.pc
}

func (t *Thread) NewFrame(method *metaspace.Method) *Frame {
	return NewFrame(method)
}

func (t *Thread)IsStackEmpty() bool {
	return t.stack.isEmpty()
}
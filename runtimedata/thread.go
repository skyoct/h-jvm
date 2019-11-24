package runtimedata

type Thread struct {
	pc    uint
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024), // 虚拟机栈的最大深度
	}
}

// 放入frame
func (t *Thread) PushFrame(frame *Frame) {
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

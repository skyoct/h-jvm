package runtimedata

// 使用链表实现栈结构

type Stack struct {
	maxSize uint // 最大数量
	size    uint // 当前的数量
	topPtr  *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{maxSize: maxSize}
}

// 栈帧入栈  top指向新的栈帧 新的栈帧的next指向原来top指向的
func (s *Stack) push(frame *Frame) {
	if s.size >= s.maxSize {
		panic("java.lang.StackOverflowError!") // 栈溢出
	}
	if s.topPtr != nil { // 如果栈不为空
		frame.next = s.topPtr
	}
	s.topPtr = frame
	s.size++
}

// 顶部栈帧
func (s *Stack) top() *Frame {
	if s.topPtr == nil {
		panic("jvm stack is empty")
	}
	return s.topPtr
}

func (s *Stack) pop() *Frame {
	if s.topPtr == nil {
		panic("jvm stack is empty")
	}
	topFrame := s.topPtr
	s.topPtr = s.topPtr.next
	s.size--
	topFrame.next = nil // 把栈帧的下一个设为nil 防止非法操作
	return topFrame
}

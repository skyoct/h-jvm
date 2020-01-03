package load

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
	"h-jvm/runtimedata/metaspace"
)

//0x2e	iaload	将int型数组指定索引的值推送至栈顶
//0x2f	laload	将long型数组指定索引的值推送至栈顶
//0×30	faload	将float型数组指定索引的值推送至栈顶
//0×31	daload	将double型数组指定索引的值推送至栈顶
//0×32	aaload	将引用型数组指定索引的值推送至栈顶
//0×33	baload	将boolean或byte型数组指定索引的值推送至栈顶
//0×34	caload	将char型数组指定索引的值推送至栈顶
//0×35	saload	将short型数组指定索引的值推送至栈顶
// 两个操作数 第一个是数组索引  第二个是数组引用

type IALoad struct {
	base.NoOperandsInstruction
}

func (i *IALoad) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	array := stack.PopRef()
	checkNotNil(array) // 判断是否空指针
	data := array.Ints()
	checkIndex(len(data), index)
	stack.PushInt(data[index]) // 压回栈顶
}

type LALoad struct {
	base.NoOperandsInstruction
}

func (L *LALoad) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	array := stack.PopRef()
	checkNotNil(array) // 判断是否空指针
	data := array.Longs()
	checkIndex(len(data), index)
	stack.PushLong(data[index]) // 压回栈顶
}


type FALoad struct {
	base.NoOperandsInstruction
}

func (F *FALoad) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	array := stack.PopRef()
	checkNotNil(array) // 判断是否空指针
	data := array.Floats()
	checkIndex(len(data), index)
	stack.PushFloat(data[index]) // 压回栈顶
}

type DALoad struct {
	base.NoOperandsInstruction
}

func (D *DALoad) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	array := stack.PopRef()
	checkNotNil(array) // 判断是否空指针
	data := array.Doubles()
	checkIndex(len(data), index)
	stack.PushDouble(data[index]) // 压回栈顶
}

type AALoad struct {
	base.NoOperandsInstruction
}

func (A *AALoad) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	array := stack.PopRef()
	checkNotNil(array) // 判断是否空指针
	data := array.Refs()
	checkIndex(len(data), index)
	stack.PushRef(data[index]) // 压回栈顶
}



// char byte/boolean s 都以int32存储


type CALoad struct {
	base.NoOperandsInstruction
}

func (C *CALoad) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	array := stack.PopRef()
	checkNotNil(array) // 判断是否空指针
	data := array.Chars()
	checkIndex(len(data), index)
	stack.PushInt(int32(data[index])) // 压回栈顶
}

type BALoad struct {
	base.NoOperandsInstruction
}

func (B *BALoad) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	array := stack.PopRef()
	checkNotNil(array) // 判断是否空指针
	data := array.Bytes()
	checkIndex(len(data), index)
	stack.PushInt(int32(data[index])) // 压回栈顶
}

type SALoad struct {
	base.NoOperandsInstruction
}

func (S *SALoad) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	array := stack.PopRef()
	checkNotNil(array) // 判断是否空指针
	data := array.Shorts()
	checkIndex(len(data), index)
	stack.PushInt(int32(data[index])) // 压回栈顶
}


func checkNotNil(ref *metaspace.Object){
	if ref == nil{
		panic("java.lang.NullPointException")
	}
}

// 检查是否越界
func checkIndex(arrLen int, index int32)  {
	if index < 0 || index >= int32(arrLen){
		panic("ArrayIndexOutOfBoundsException")
	}
}
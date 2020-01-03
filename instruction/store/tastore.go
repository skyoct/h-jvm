package store

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
	"h-jvm/runtimedata/metaspace"
)

// 0x4f	iastore	将栈顶int型数值存入指定数组的指定索引位置
//0×50	lastore	将栈顶long型数值存入指定数组的指定索引位置
//0×51	fastore	将栈顶float型数值存入指定数组的指定索引位置
//0×52	dastore	将栈顶double型数值存入指定数组的指定索引位置
//0×53	aastore	将栈顶引用型数值存入指定数组的指定索引位置
//0×54	bastore	将栈顶boolean或byte型数值存入指定数组的指定索引位置
//0×55	castore	将栈顶char型数值存入指定数组的指定索引位置
//0×56	sastore	将栈顶short型数值存入指定数组的指定索引位置

// 需要三个操作数 分别是：要赋值给元素的值，数组索引 数组引用
// 从操作数栈中依次弹出

type IAStore struct {
	base.NoOperandsInstruction
}

func (i * IAStore) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	array := stack.PopRef()
	checkNotNil(array)
	data := array.Ints()
	checkIndex(len(data), index)
	data[index] = int32(val)
}

type LAStore struct {
	base.NoOperandsInstruction
}

func (L * LAStore) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	index := stack.PopInt()
	array := stack.PopRef()
	checkNotNil(array)
	data := array.Longs()
	checkIndex(len(data), index)
	data[index] = int64(val)
}

type FAStore struct {
	base.NoOperandsInstruction
}

func (F * FAStore) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	index := stack.PopInt()
	array := stack.PopRef()
	checkNotNil(array)
	data := array.Floats()
	checkIndex(len(data), index)
	data[index] = val
}

type DAStore struct {
	base.NoOperandsInstruction
}

func (D * DAStore) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	index := stack.PopInt()
	array := stack.PopRef()
	checkNotNil(array)
	data := array.Doubles()
	checkIndex(len(data), index)
	data[index] = val
}

type AAStore struct {
	base.NoOperandsInstruction
}

func (A * AAStore) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopRef()
	index := stack.PopInt()
	array := stack.PopRef()
	checkNotNil(array)
	data := array.Refs()
	checkIndex(len(data), index)
	data[index] = val
}

type CAStore struct {
	base.NoOperandsInstruction
}

func (C * CAStore) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	array := stack.PopRef()
	checkNotNil(array)
	data := array.Chars()
	checkIndex(len(data), index)
	data[index] = uint16(val)
}

type BAStore struct {
	base.NoOperandsInstruction
}

func (B * BAStore) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	array := stack.PopRef()
	checkNotNil(array)
	data := array.Bytes()
	checkIndex(len(data), index)
	data[index] = int8(val)
}

type SAStore struct {
	base.NoOperandsInstruction
}

func (S * SAStore) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	array := stack.PopRef()
	checkNotNil(array)
	data := array.Shorts()
	checkIndex(len(data), index)
	data[index] = int16(val)
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


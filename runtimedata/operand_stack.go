package runtimedata

import "math"

// 操作数栈

type OperandStack struct {
	size  uint // 当前元素个数 其实指向栈顶元素的上一个
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

func (o *OperandStack) PushInt(val int32) {
	o.slots[o.size].val = val
	o.size++
}

func (o *OperandStack) PopInt() int32 {
	o.size--
	return o.slots[o.size].val
}

func (o *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	o.slots[o.size].val = int32(bits)
	o.size++
}

func (o *OperandStack) PopFloat() float32 {
	o.size--
	bits := uint32(o.slots[o.size].val)
	return math.Float32frombits(bits)
}

func (o *OperandStack) PushLong(val int64) {
	o.slots[o.size].val = int32(val) // 存放高位-》存放低位 大端序
	o.slots[o.size+1].val = int32(val >> 32)
	o.size += 2
}

func (o *OperandStack) PopLong() int64 {
	o.size -= 2
	lower := o.slots[o.size].val          // 获取低位的
	high := o.slots[o.size+1].val         //获取高位
	return int64(high)<<32 | int64(lower) // 高位<<32 低位补0 和lower | 相当于相加
}

func (o *OperandStack) PushDouble(val float64) {
	longVal := int64(math.Float64bits(val))
	o.PushLong(longVal)
}

func (o *OperandStack) PopDouble() float64 {
	bits := uint64(o.PopLong())
	return math.Float64frombits(bits)
}

func (o *OperandStack) PushRef(ref *Reference) {
	o.slots[o.size].ref = ref
	o.size++
}

func (o *OperandStack) PopRef() *Reference {
	o.size--
	return o.slots[o.size].ref
}

func (o *OperandStack) PopSlot() {
	o.size--
}

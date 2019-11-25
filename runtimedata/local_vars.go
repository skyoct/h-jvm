package runtimedata

import (
	"math"
)

// 一个局部变量表包含多个slot
type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (l LocalVars) SetInt(index uint, val int32) {
	l[index].val = val
}

func (l LocalVars) GetInt(index uint) int32 {
	return l[index].val
}

// 把float转化为二进制，然后二进制以int存储
func (l LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	l[index].val = int32(bits)
}

// 先把int32转化为byte 然后转换为浮点型
func (l LocalVars) GetFloat(index uint) float32 {
	bits := uint32(l[index].val)
	return math.Float32frombits(bits)
}

// 长整型需要拆分2个 存储在两个int32
func (l LocalVars) SetLong(index uint, val int64) {
	l[index].val = int32(val & 0xFFFFFFFF) // 转换为32 丢失高位
	l[index+1].val = int32(val >> 32)      // 丢失低位  把高位向右移动八位
}

// 把存储在两个int32的拼接成为一个64位的
func (l LocalVars) GetLong(index uint) int64 {
	lower := l[index].val                 // 获取低位的
	high := l[index+1].val                //获取高位
	return int64(high)<<32 | int64(lower) // 高位<<32低位补0 和lower | 相当于相加
}

// double转换为long类型 然后转换为两个int存储
func (l LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	l.SetLong(index, int64(bits))
}

// long转换为double
func (l LocalVars) GetDouble(index uint) float64 {
	return float64(math.Float64frombits(uint64(l.GetLong(index))))
}

// 设置引用
func (l LocalVars) SetRef(index uint, ref *Reference) {
	l[index].ref = ref
}

func (l LocalVars) GetRef(index uint) *Reference {
	return l[index].ref
}

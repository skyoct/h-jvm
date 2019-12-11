package metaspace

import "math"

//  用来存放静态变量和实例变量
type Slot struct {
	val int32      // 存放int float （double和lang需要两个）
	ref *Reference // 存放引用
}

type Slots []Slot

func newSlots(slotCount uint) Slots {
	if slotCount > 0 {
		return make([]Slot, slotCount)
	}
	return nil
}

func (s Slots) SetInt(index uint, val int32) {
	s[index].val = val
}
func (s Slots) GetInt(index uint) int32 {
	return s[index].val
}

func (s Slots) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	s[index].val = int32(bits)
}
func (s Slots) GetFloat(index uint) float32 {
	bits := uint32(s[index].val)
	return math.Float32frombits(bits)
}

// long consumes two slots
func (s Slots) SetLong(index uint, val int64) {
	s[index].val = int32(val)
	s[index+1].val = int32(val >> 32)
}
func (s Slots) GetLong(index uint) int64 {
	low := uint32(s[index].val)
	high := uint32(s[index+1].val)
	return int64(high)<<32 | int64(low)
}

// double consumes two slots
func (s Slots) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	s.SetLong(index, int64(bits))
}
func (s Slots) GetDouble(index uint) float64 {
	bits := uint64(s.GetLong(index))
	return math.Float64frombits(bits)
}

func (s Slots) SetRef(index uint, ref *Reference) {
	s[index].ref = ref
}
func (s Slots) GetRef(index uint) *Reference {
	return s[index].ref
}

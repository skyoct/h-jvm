package metaspace


func (o *Object) Bytes() []int8 {
	return o.data.([]int8)
}
func (o *Object) Shorts() []int16 {
	return o.data.([]int16)
}
func (o *Object) Ints() []int32 {
	return o.data.([]int32)
}
func (o *Object) Longs() []int64 {
	return o.data.([]int64)
}
func (o *Object) Chars() []uint16 {
	return o.data.([]uint16)
}
func (o *Object) Floats() []float32 {
	return o.data.([]float32)
}
func (o *Object) Doubles() []float64 {
	return o.data.([]float64)
}
func (o *Object) Refs() []*Object {
	return o.data.([]*Object)
}

func (o *Object) ArrayLength() int32 {
	switch o.data.(type) {
	case []int8:
		return int32(len(o.data.([]int8)))
	case []int16:
		return int32(len(o.data.([]int16)))
	case []int32:
		return int32(len(o.data.([]int32)))
	case []int64:
		return int32(len(o.data.([]int64)))
	case []uint16:
		return int32(len(o.data.([]uint16)))
	case []float32:
		return int32(len(o.data.([]float32)))
	case []float64:
		return int32(len(o.data.([]float64)))
	case []*Object:
		return int32(len(o.data.([]*Object)))
	default:
		panic("Not array!")
	}
}
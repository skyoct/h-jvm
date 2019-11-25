package base

// 解析读取code表中的指令
type CodeReader struct {
	code []byte
	pc   int
}

func (c *CodeReader) Reset(code []byte, pc int) {
	c.code = code
	c.pc = pc
}

func (c *CodeReader) ReadUint8() uint8 {
	i := c.code[c.pc]
	c.pc++ // 程序计数器加1
	return i
}

func (c *CodeReader) ReadInt8() int8 {
	return int8(c.ReadUint8())
}

func (c *CodeReader) ReadUint16() uint16 {
	b1 := c.ReadUint8()
	b2 := c.ReadUint8()
	return uint16(b1)<<8 | uint16(b2) // 先转换为uint16 防止左移溢出
}

func (c *CodeReader) ReadInt16() int16 {
	return int16(c.ReadUint16())
}

func (c *CodeReader) ReadInt32() int32 {
	b1 := c.ReadUint8()
	b2 := c.ReadUint8()
	b3 := c.ReadUint8()
	b4 := c.ReadUint8()
	return int32(b1)<<24 | int32(b2)<<16 | int32(b3)<<8 | int32(b4)
}

// 有0-3的padding 保证获取到的偏移是4的倍数
func (c *CodeReader) SkipPadding() {
	for c.pc%4 != 0 {
		c.ReadUint8()
	}
}

// Switch 使用
func (c *CodeReader) ReadInt32s(n int32) []int32 {
	ints := make([]int32, n)
	for i := range ints {
		ints[i] = c.ReadInt32()
	}
	return ints
}

func (c *CodeReader) Pc() int {
	return c.pc
}

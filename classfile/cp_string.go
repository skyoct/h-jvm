package classfile


// 字符类型常量
type ConstantStringInfo struct {
	cp ConstantPool
	index uint16
}

func (c *ConstantStringInfo) readInfo(reader *ClassReader) {
	c.index = reader.readUint16()
}

func (c *ConstantStringInfo)String() string {
	return c.cp.getUtf8(c.index)
}





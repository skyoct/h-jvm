package classfile


// 类类型常量
type ConstantClassInfo struct {
	cp ConstantPool
	index uint16
}

func (c *ConstantClassInfo) readInfo(reader *ClassReader) {
	c.index = reader.readUint16()
}

func (c *ConstantClassInfo) Name() string {
	return c.cp.getClassName(c.index)
}





package classfile


// ConstantValue是定长属性，只会出现在field_info结构中，用于 表示常量表达式的值
type ConstantValueAttribute struct {
	index uint16
}

func (c *ConstantValueAttribute) readInfo(reader *ClassReader) {
	c.index = reader.readUint16()
}

func (c *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return c.index
}




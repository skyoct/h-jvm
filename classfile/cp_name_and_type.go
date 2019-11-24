package classfile

// 指向字段或者方法的名称和描述符
type ConstantNameAndTypeInfo struct {
	nameIndex uint16 // 指向该字段或者方法的名称常量项的索引
	descIndex uint16 // 指向该字段或者方法的描述符常量项的索引
}

func (c *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	c.nameIndex = reader.readUint16()
	c.descIndex = reader.readUint16()
}

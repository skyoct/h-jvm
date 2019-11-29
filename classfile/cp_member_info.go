package classfile

// 抽取Fieldref Methodref interfaceMethodref的共同结构
type ConstantMemberInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (c *ConstantMemberInfo) ClassName() string {
	return c.cp.getClassName(c.classIndex)
}

func (c *ConstantMemberInfo) NameAndDescriptor() (string, string) {
	return c.cp.getNameAndType(c.nameAndTypeIndex)
}

func (c *ConstantMemberInfo) readInfo(reader *ClassReader) {
	c.classIndex = reader.readUint16()
	c.nameAndTypeIndex = reader.readUint16()
}

//CONSTANT_Fieldref_info表示字段符号引用
type ConstantFieldrefInfo struct {
	ConstantMemberInfo
}

// CONSTANT_Methodref_info表示普通（非接口）方法符号引用，
type ConstantMethodrefInfo struct {
	ConstantMemberInfo
}

// CONSTANT_InterfaceMethodref_info表示接口方法符号引用
type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberInfo
}

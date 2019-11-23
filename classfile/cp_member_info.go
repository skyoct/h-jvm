package classfile


//CONSTANT_Fieldref_info表示字段符号引用
type ConstantFieldrefInfo struct {
	cp ConstantPool
	classIndex uint16
	nameAndTypeIndex uint16
}
func (c *ConstantFieldrefInfo) readInfo(reader *ClassReader) {
	c.classIndex = reader.readUint16()
	c.nameAndTypeIndex = reader.readUint16()
}

func (c *ConstantFieldrefInfo) ClassName() string {
	return c.cp.getClassName(c.classIndex)
}

func (c *ConstantFieldrefInfo) NameAndDescriptor() (string, string) {
	return c.cp.getNameAndType(c.nameAndTypeIndex)
}


// CONSTANT_Methodref_info表示普通（非接口）方法符号引用，
type ConstantMethodrefInfo struct {
	cp ConstantPool
	classIndex uint16
	nameAndTypeIndex uint16
}
func (c *ConstantMethodrefInfo) readInfo(reader *ClassReader) {
	c.classIndex = reader.readUint16()
	c.nameAndTypeIndex = reader.readUint16()
}

func (c *ConstantMethodrefInfo) ClassName() string {
	return c.cp.getClassName(c.classIndex)
}

func (c *ConstantMethodrefInfo) NameAndDescriptor() (string, string) {
	return c.cp.getNameAndType(c.nameAndTypeIndex)
}

// CONSTANT_InterfaceMethodref_info表示接口方法符号引用
type ConstantInterfaceMethodrefInfo struct {
	cp ConstantPool
	classIndex uint16
	nameAndTypeIndex uint16
}
func (c *ConstantInterfaceMethodrefInfo) readInfo(reader *ClassReader) {
	c.classIndex = reader.readUint16()
	c.nameAndTypeIndex = reader.readUint16()
}

func (c *ConstantInterfaceMethodrefInfo) ClassName() string {
	return c.cp.getClassName(c.classIndex)
}

func (c *ConstantInterfaceMethodrefInfo) NameAndDescriptor() (string, string) {
	return c.cp.getNameAndType(c.nameAndTypeIndex)
}

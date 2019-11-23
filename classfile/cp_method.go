package classfile

/*
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
*/
type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (c *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	c.referenceKind = reader.readUint8()
	c.referenceIndex = reader.readUint16()
}

/*
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
*/
type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

func (c *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	c.descriptorIndex = reader.readUint16()
}

/*
CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
*/
type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (c *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	c.bootstrapMethodAttrIndex = reader.readUint16()
	c.nameAndTypeIndex = reader.readUint16()
}

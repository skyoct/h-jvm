package classfile

/*
field_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
method_info {
    u2             access_flags;
    u2             name_index;
    u2             descriptor_index;
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/

// 字段和方法的结构体

type MemberInfo struct {
	cp ConstantPool  // 常量池指针
	accessFlags uint16 // 访问标志
	nameIndex uint16 // 名字索引
	descIndex uint16 // 描述符索引
	attributes []AttributeInfo
}


// 读取所有的方法或者字段
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

// 读取一个方法字段
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp: cp,
		accessFlags: reader.readUint16(),
		nameIndex: reader.readUint16(),
		descIndex: reader.readUint16(),
		attributes: readAttributes(reader, cp),
	}
}

func (m *MemberInfo) AccessFlags() uint16 {
	return m.accessFlags
}
func (m *MemberInfo) Name() string {
	return m.cp.getUtf8(m.nameIndex)
}
func (m *MemberInfo) Descriptor() string {
	return m.cp.getUtf8(m.descIndex)
}

func (m *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range m.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

func (m *MemberInfo) ConstantValueAttribute() *ConstantValueAttribute {
	for _, attrInfo := range m.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}

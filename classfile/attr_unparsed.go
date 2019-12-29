package classfile

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (u *UnparsedAttribute) readInfo(reader *ClassReader) {
	u.info = reader.readBytes(u.length)
}

func (u *UnparsedAttribute) Info() []byte {
	return u.info
}

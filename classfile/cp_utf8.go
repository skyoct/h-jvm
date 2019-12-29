package classfile

// utf8类型常量

/**
CONSTANT_Utf8_info {
	u1 tag;
	u2 length;
	u1 bytes[length];
}
*/

type ConstantUtf8Info struct {
	val string
}

func (c *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16()) // 长度
	bytes := reader.readBytes(length)     // 读取
	c.val = decodeMUTF8(bytes)
}

func (c *ConstantUtf8Info) Str() string {
	return c.val
}

func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}

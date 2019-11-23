package classfile

import "encoding/binary"

/**
	来读取字节码文件

 */

type ClassReader struct {
	data[] byte
}

// 读取一个字节
func (c *ClassReader) readUint8() uint8 {
	val := c.data[0]
	c.data = c.data[1:]
	return val
}
// 读取两个字节
func (c *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(c.data)
	c.data = c.data[2:]
	return val
}
// 读取四个字节
func (c *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(c.data)
	c.data = c.data[4:]
	return val
}
// 读取八个字节
func (c *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(c.data)
	c.data = c.data[8:]
	return val
}

// 读取两个字节的数组 前两个字节为读取的长度
func (c *ClassReader) readUint16s() []uint16 {
	n := c.readUint16()
	//c.data = c.data[2:]
	s := make([]uint16, n)
	for i := range s {
		s[i] = c.readUint16()
	}
	return s
}

// 读取字节数组
func (c *ClassReader) readBytes(n uint32) []byte {
	bytes := c.data[ :n]
	c.data = c.data[n: ]
	return bytes
}



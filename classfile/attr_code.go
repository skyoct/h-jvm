package classfile

type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	Attributes     []AttributeInfo
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (c *CodeAttribute) readInfo(reader *ClassReader) {
	c.maxStack = reader.readUint16()
	c.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	c.code = reader.readBytes(codeLength)
	c.exceptionTable = readExceptionTable(reader)
	c.Attributes = readAttributes(reader, c.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}

func (c *CodeAttribute) Code() []byte {
	return c.code
}

func (c *CodeAttribute) MaxLocals() uint {
	return uint(c.maxLocals)
}

func (c *CodeAttribute) MaxStack() uint {
	return uint(c.maxStack)
}

package classfile

type ExceptionAttribute struct {
	indexTable []uint16
}

func (e *ExceptionAttribute) readInfo(reader *ClassReader) {
	e.indexTable = reader.readUint16s()
}

func (e *ExceptionAttribute) ExceptionIndexTable() []uint16 {
	return e.indexTable
}




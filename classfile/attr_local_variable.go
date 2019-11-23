package classfile

// 描述栈帧中局部变量表中的变量与java源码中对于的关系
type LocalVariableTableAttribute struct {
	LocalVariableTable []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc uint16
	length uint16
	nameIndex uint16
	descIndex uint16
	index uint16
}

func (l *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	localVariableTable := reader.readUint16()
	l.LocalVariableTable = make([]*LocalVariableTableEntry, localVariableTable)
	for i := range l.LocalVariableTable {
		l.LocalVariableTable[i] = &LocalVariableTableEntry{
			startPc: reader.readUint16(),
			length: reader.readUint16(),
			nameIndex: reader.readUint16(),
			descIndex: reader.readUint16(),
			index: reader.readUint16(),
		}
	}
}

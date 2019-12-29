package classfile


// SourceFile是可选定长属性，只会出现在ClassFile结构中，用于指出源文件
type SourceFileAttribute struct {
	cp ConstantPool
	index uint16
}

func (s *SourceFileAttribute) readInfo(reader *ClassReader) {
	s.index = reader.readUint16()
}

func (s *SourceFileAttribute) FileName() string {
	return s.cp.getUtf8(s.index)
}



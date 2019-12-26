package classfile

//
//StackMapTable_attribute {
//u2              attribute_name_index;
//u4              attribute_length;
//u2              number_of_entries;
//stack_map_frame entries[number_of_entries];
//}

// StackMapTable主要用来验证跳转前后locals、stack中的类型和大小一致。
// 当前先跳过实现
// 具体实现参考 Java虚拟机规范.Java SE 8版 87页

type StackMapTable struct {
	len uint32
}

func (s *StackMapTable) readInfo(reader *ClassReader) {
	reader.readBytes(s.len)
}

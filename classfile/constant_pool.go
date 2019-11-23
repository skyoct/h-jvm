package classfile

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool{
	cpCount := int(reader.readUint16()) // 读取常量池大小 使用u2类型的数据表示 计数容量从1开始
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i<cpCount; i++ {  // 索引从1开始 0被空出来
		cp[i] = ReadConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo: // 占两个位置
			i++
		}
	}
	return cp
}

// 根据索引查找常量信息
func (c ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := c[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

// 根据索引从常量池中找出utf-8 字符串
func (c ConstantPool) getUtf8(index uint16) string {
	utf8Info := c.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.val
}

// 从常量池中查找类名称
func (c ConstantPool) getClassName(index uint16) string {
	classInfo := c.getConstantInfo(index).(*ConstantClassInfo)
	return c.getUtf8(classInfo.index)
}


// 从常量池中找出方法或者字段的描述符
func (c ConstantPool) getNameAndType(index uint16) (string, string){
	ntInfo := c.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := c.getUtf8(ntInfo.nameIndex)
	_type := c.getUtf8(ntInfo.descIndex)
	return name, _type
}




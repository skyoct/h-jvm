package classfile

const (
	CONSTANT_utf8 = 1
	CONSTANT_Integer = 3
	CONSTANT_Float = 4
	CONSTANT_Long = 5
	CONSTANT_Double = 6
	CONSTANT_Class = 7
	CONSTANT_String = 8
	CONSTANT_Fieldref = 9
	CONSTANT_Methodref = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_NameAndType = 12
	CONSTANT_MethodHandle = 15
	CONSTANT_MethodType = 16
	CONSTANT_InvokeDynamic = 17
)

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func ReadConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8() // 读取tag  然后判断是什么类型的常量
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_Class:
		return &ConstantClassInfo{cp:cp}
	case CONSTANT_String:
		return &ConstantStringInfo{cp:cp}
	case CONSTANT_Fieldref:
		return &ConstantFieldrefInfo{cp:cp}
	case CONSTANT_Methodref:
		return &ConstantMethodrefInfo{cp:cp}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{cp:cp}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("")
	}
}



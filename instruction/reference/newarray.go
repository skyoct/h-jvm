package reference

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
	"h-jvm/runtimedata/metaspace"
)
// 0xbc
// newarray是创建基本类型的数组
// newarray指令后面跟着两个操作数，第一个操作数是uint8整形，表示要创建那种类型的值, java虚拟机规范把这个值叫做ATYPE
// 第二个操作数是count 从操作数栈弹出

const (
	ATYPE_BOOLEAN = 4
	ATYPE_CHAR    = 5
	ATYPE_FLOAT   = 6
	ATYPE_DOUBLE  = 7
	ATYPE_BYTE    = 8
	ATYPE_SHORT   = 9
	ATYPE_INT     = 10
	ATYPE_LONG    = 11
)

type NewArray struct {
	atype uint8
}

func (n *NewArray) FetchOperands(reader *base.CodeReader) {
	n.atype = reader.ReadUint8()
}

func (n *NewArray) Execute(frame *runtimedata.Frame) {
	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0{
		panic("java.lang.NegativeArraySizeException")
	}
	classLoader := frame.Method().Class().Loader() // 获取类加载器
	arrClass := getPrimeArrayClass(classLoader, n.atype)
	array := arrClass.NewArray(uint(count))
	stack.PushRef(array) //  押入操作数栈
}

func getPrimeArrayClass(loader *metaspace.ClassLoader, atype uint8) *metaspace.Class{
	switch atype {
	case ATYPE_BOOLEAN:
		return loader.LoadClass("[Z")
	case ATYPE_BYTE:
		return loader.LoadClass("[B")
	case ATYPE_CHAR:
		return loader.LoadClass("[C")
	case ATYPE_SHORT:
		return loader.LoadClass("[S")
	case ATYPE_INT:
		return loader.LoadClass("[I")
	case ATYPE_LONG:
		return loader.LoadClass("[J")
	case ATYPE_FLOAT:
		return loader.LoadClass("[F")
	case ATYPE_DOUBLE:
		return loader.LoadClass("[D")
	default:
		panic("Invalid atype!")
	}
}



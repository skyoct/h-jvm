package classfile

import "fmt"

// 属性表

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "BootstrapMethods":
		return &BootstrapMethodsAttribute{}
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionAttribute{}
	case "InnerClasses":
		return &InnerClassesAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	case "Signature":
		return &SignatureAttribute{cp: cp}
	case "StackMapTable":
		return &StackMapTable{attrLen}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		fmt.Println(attrName)
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}

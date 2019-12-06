package metaspace

import "h-jvm/classfile"

type MemberRef struct {
	SymRef
	name       string // 名字
	descriptor string // 描述符
}

func (c *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberInfo) {
	c.className = refInfo.ClassName()
	c.name, c.descriptor = refInfo.NameAndDescriptor()
}

// 字段符号
type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberInfo)
	return ref
}

// 方法符号引用
type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberInfo)
	return ref
}

// 接口方法引用
type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberInfo)
	return ref
}

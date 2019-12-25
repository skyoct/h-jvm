package metaspace

import "h-jvm/classfile"

type Method struct {
	ClassMember
	maxStack  uint
	maxLocals uint
	code      []byte
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttribute(cfMethod)
	}
	return methods
}

func (m *Method) copyAttribute(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		m.maxLocals = codeAttr.MaxLocals()
		m.maxStack = codeAttr.MaxStack()
		m.code = codeAttr.Code()
	}
}

// getter

func (m *Method) MaxLocals() uint {
	return m.maxLocals
}

func (m *Method) MaxStack() uint {
	return m.maxStack
}

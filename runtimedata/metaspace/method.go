package metaspace

import "h-jvm/classfile"

type Method struct {
	ClassMember
	maxStack     uint
	maxLocals    uint
	code         []byte
	argSlotCount uint
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = &Method{}
		methods[i].class = class
		methods[i].copyMemberInfo(cfMethod)
		methods[i].copyAttribute(cfMethod)
		methods[i].calcArgSlotCount()
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

func (m *Method) calcArgSlotCount() {
	parsedDescriptor := parseMethodDescriptor(m.descriptor)
	for _, paramType := range parsedDescriptor.parameterTypes {
		m.argSlotCount++
		if paramType == "J" || paramType == "D" {
			m.argSlotCount++
		}
	}
	if !m.IsStatic() {
		m.argSlotCount++ // `this` reference
	}
}

// 修饰符
func (m *Method) IsSynchronized() bool {
	return 0 != m.accessFlags&ACC_SYNCHRONIZED
}
func (m *Method) IsBridge() bool {
	return 0 != m.accessFlags&ACC_BRIDGE
}
func (m *Method) IsVarargs() bool {
	return 0 != m.accessFlags&ACC_VARARGS
}
func (m *Method) IsNative() bool {
	return 0 != m.accessFlags&ACC_NATIVE
}
func (m *Method) IsAbstract() bool {
	return 0 != m.accessFlags&ACC_ABSTRACT
}
func (m *Method) IsStrict() bool {
	return 0 != m.accessFlags&ACC_STRICT
}

// getter

func (m *Method) MaxLocals() uint {
	return m.maxLocals
}

func (m *Method) MaxStack() uint {
	return m.maxStack
}

func (m *Method) Code() []byte {
	return m.code
}

func (m *Method) ArgsSlotCount() uint {
	return m.argSlotCount
}

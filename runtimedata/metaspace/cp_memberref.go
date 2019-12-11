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

//字段的符号引用解析 如果存在 直接返回 不存在就去解析
func (f *FieldRef) ResolvedField() *Field {
	if f.field == nil {

	}
	return f.field
}

// 查找字段 。如果字段查找失败，则虚拟机抛出NoSuchFieldError异常。
// 如果查找成功， 但D没有足够的权限访问该字段，则虚拟机抛出IllegalAccessError异常
func (f *FieldRef) resolveFieldRef() {
	d := f.cp.class
	c := f.ResolvedClass() // 先解决类的符号引用
	field := lookupField(c, f.name, f.descriptor)
	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	f.field = field
}

// 现在当前类查找 没找到去接口和父类查找
func lookupField(c *Class, name, descriptor string) *Field {
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor { // 如果名字和描述符相同
			return field
		}
	}

	for _, superInterface := range c.interfaces {
		lookupField(superInterface, name, descriptor) // 去接口上查找
	}

	if c.superClass != nil {
		lookupField(c.superClass, name, descriptor)
	}
	return nil
}

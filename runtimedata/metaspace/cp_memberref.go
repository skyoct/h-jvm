package metaspace

import (
	"h-jvm/classfile"
)

type MemberRef struct {
	SymRef
	name       string // 名字
	descriptor string // 描述符
}

func (m *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberInfo) {
	m.className = refInfo.ClassName()
	m.name, m.descriptor = refInfo.NameAndDescriptor()
}

func (m *MemberRef) Descriptor() string {
	return m.descriptor
}

func (m *MemberRef) Name() string {
	return m.name
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

//字段的符号引用解析 如果存在 直接返回 不存在就去解析
func (f *FieldRef) ResolvedField() *Field {
	if f.field == nil {
		f.resolveFieldRef()
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
	if !field.isAccessibleTo(d) { // 判断是否有访问权限
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

// 方法符号引用
type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberInfo)
	return ref
}

func (m *MethodRef) ResolvedMethod() *Method {
	if m.method == nil {
		m.resolveMethodRef()
	}
	return m.method
}

// 解决方法引用
func (m *MethodRef) resolveMethodRef() {
	class := m.cp.class // 获取当前方法在哪个类内
	currentClass := m.ResolvedClass()
	if currentClass.IsInterface() { // 判断要执行的方法的类是不是接口
		panic("java.lang.IncompatibleClassChangeError")
	}
	method := lookupMethod(currentClass, m.name, m.descriptor)
	//fmt.Println("*********")
	//fmt.Println(currentClass.Name())
	//fmt.Println(m.descriptor)
	//fmt.Println("*********", m.name)

	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(class) { // 判断是否有访问权限
		panic("java.lang.IllegalAccessError")
	}
	m.method = method
}

// 查找方法
func lookupMethod(class *Class, name, descriptor string) *Method {
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}

// 接口方法引用
type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(cp *ConstantPool, refInfo *classfile.ConstantInterfaceMethodrefInfo) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberInfo)
	return ref
}

func (i *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if i.method == nil {

	}
	return i.method
}

func (i *InterfaceMethodRef) resolveInterfaceMethodRef() {
	d := i.cp.class
	c := i.ResolvedClass() // 解决类的加载
	if !c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	method := lookupInterfaceMethod(c, i.name, i.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	i.method = method
}

func lookupInterfaceMethod(ifaces *Class, name, descriptor string) *Method {
	for _, method := range ifaces.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return lookupMethodInInterfaces(ifaces.interfaces, name, descriptor)
}

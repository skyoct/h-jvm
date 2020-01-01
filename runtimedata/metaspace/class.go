package metaspace

import (
	"h-jvm/classfile"
	"strings"
)

type Class struct {
	accessFlags       uint16 // 访问标志
	name              string // class name
	superClassName    string
	interfaceNames    []string      // 接口名字集合
	constantPool      *ConstantPool // 常量池
	fields            []*Field
	methods           []*Method
	classLoader       *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint  // 实例变量槽数量
	staticSlotCount   uint  // 静态变量槽数量
	staticVars        Slots // 静态变量放的槽

}

// 创建类
func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newField(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

// 创建一个对象
func (c *Class) NewObject() *Object {
	return newObject(c)
}

// 判断标志文件
func (c *Class) IsPublic() bool {
	return 0 != c.accessFlags&ACC_PUBLIC
}
func (c *Class) IsFinal() bool {
	return 0 != c.accessFlags&ACC_FINAL
}
func (c *Class) IsSuper() bool {
	return 0 != c.accessFlags&ACC_SUPER
}
func (c *Class) IsInterface() bool {
	return 0 != c.accessFlags&ACC_INTERFACE
}
func (c *Class) IsAbstract() bool {
	return 0 != c.accessFlags&ACC_ABSTRACT
}
func (c *Class) IsSynthetic() bool {
	return 0 != c.accessFlags&ACC_SYNTHETIC
}
func (c *Class) IsAnnotation() bool {
	return 0 != c.accessFlags&ACC_ANNOTATION
}
func (c *Class) IsEnum() bool {
	return 0 != c.accessFlags&ACC_ENUM
}

// getters
func (c *Class) ConstantPool() *ConstantPool {
	return c.constantPool
}
func (c *Class) StaticVars() Slots {
	return c.staticVars
}

// 判断是否能访问另外一个类
func (c *Class) isAccessibleTo(other *Class) bool {
	return c.IsPublic() ||
		c.getPackageName() == other.getPackageName()
}

// 得到包名
func (c *Class) getPackageName() string {
	if i := strings.LastIndex(c.name, "/"); i >= 0 {
		return c.name[:i]
	}
	return ""
}

// 得到main函数
func (c *Class) GetMainMethod() *Method {
	return c.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (c *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range c.methods { // 循环方法列表 找到静态标志 名字 方法的描述符相同的方法
		if method.IsStatic() &&
			method.name == name &&
			method.descriptor == descriptor {

			return method
		}
	}
	return nil
}

func (c *Class) Name() string {
	return c.name
}

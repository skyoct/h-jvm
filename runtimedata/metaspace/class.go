package metaspace

import "h-jvm/classfile"

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

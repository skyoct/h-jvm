package metaspace

import (
	"fmt"
	"h-jvm/classfile"
	"h-jvm/classpath"
)

type ClassLoader struct {
	cp       *classpath.Classpath
	classMap map[string]*Class //key是类的全限定名 value是类
}

func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		cp:       cp,
		classMap: make(map[string]*Class),
	}
}

// 加载类 如果类已经存在则直接返回类 不存在则加载类
func (c *ClassLoader) LoadClass(name string) *Class {
	if class, ok := c.classMap[name]; ok {  //类已经被加载
		return class
	}
	if name[0] == '[' {   //如果为数组
		//return c.\
		return c.loadArrayClass(name)

	}
	return c.loadNonArrayClass(name)
}

// 加载不是数组的类
func (c *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := c.readClass(name)
	class := c.defineClass(data)
	link(class)
	fmt.Printf("[Loaded %s from %s]\n", name, entry)
	return class
}

// 加载是数组的类
// 父类是Object
// 实现了 Cloneable 和 Serializable接口
func (c *ClassLoader) loadArrayClass(name string) *Class{
	class := &Class{
		accessFlags: ACC_PUBLIC,
		name: name,
		classLoader: c,
		initStarted: true,
		superClass: c.LoadClass("java/lang/Object"),
		interfaces: []*Class{
			c.LoadClass("java/lang/Cloneable"),
			c.LoadClass("java/io/Serializable"),
		},
	}
	c.classMap[name] = class // 添加到数组中
	return class
}

func (c *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := c.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

func (c *ClassLoader) defineClass(data []byte) *Class {
	class := parserClass(data)
	class.classLoader = c
	// 解决父类
	resolveSuperClass(class)
	// 解决接口
	resolveInterfaces(class)
	// 在classloader里面加上这个类
	class.classLoader.classMap[class.name] = class
	return class
}

func parserClass(data []byte) *Class {
	cf, err := classfile.Parser(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}

// 主要来加载父类
func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" { // 在java中只有Object类没有父类
		class.superClass = class.classLoader.LoadClass(class.superClassName) // 加载父类
	}
}

// 加载接口
func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)    // 创建数组
		for i, interfaceName := range class.interfaceNames { // 循环加载接口
			class.interfaces[i] = class.classLoader.LoadClass(interfaceName)
		}
	}
}

// 连接（link） 包括验证 准备 解析三个阶段
func link(class *Class) {
	verify()
	prepare(class)
}

// 准备阶段
func verify() {

}

//准备阶段是为类变量分配内存并且设置初始化值的阶段，这些变量所使用的内存都在元空间分配
// 这个阶段初始化的数据只有静态字段，并且是赋值初始化值（final字段除外）
func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStatic(class)
}

// 计算实例变量所占空间
func calcInstanceFieldSlotIds(class *Class) {
	{
		slotId := uint(0)
		if class.superClass != nil {
			slotId = class.superClass.instanceSlotCount // 父类的实例变量槽数量
		}
		for _, field := range class.fields {
			if !field.IsStatic() { // 不是静态字段
				field.slotId = slotId
				slotId++
				if field.isDoubleOrLong() { // 如果为double或者long占用两个slot
					slotId++
				}
			}
		}
		class.instanceSlotCount = slotId
	}
}

// 计算静态变量变量所占空间
func calcStaticFieldSlotIds(class *Class) {
	{
		slotId := uint(0)
		for _, field := range class.fields {
			if field.IsStatic() { // 是静态字段
				field.slotId = slotId
				slotId++
				if field.isDoubleOrLong() { // 如果为double或者long占用两个slot
					slotId++
				}
			}
		}
		class.staticSlotCount = slotId
	}
}

// 为静态变量分配空间 并赋予初始值 赋予初始值 (final直接赋予给定的值)
func allocAndInitStatic(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount) // 分配空间
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field) //初始化final类型静态变量的值
		}
	}
}

// 初始化final静态变量
func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.slotId
	slotId := field.slotId
	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32) // 使用断言转
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64) // 使用断言转
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32) // 使用断言转
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64) // 使用断言转
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			// 等待实现
		}
	}
}

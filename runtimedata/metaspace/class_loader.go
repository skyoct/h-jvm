package metaspace

import (
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
	if class, ok := c.classMap[name]; ok {
		return class
	}
	return
}

func (c *ClassLoader) loadNonArrayClass(name string) {
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
	resolveIntefaces(class)
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
func resolveIntefaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)    // 创建数组
		for i, interfaceName := range class.interfaceNames { // 循环加载接口
			class.interfaces[i] = class.classLoader.LoadClass(interfaceName)
		}
	}
}

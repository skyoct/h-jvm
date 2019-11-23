package classfile

import "fmt"

type ClassFile struct {
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags uint16
	thisClass uint16
	superClass uint16
	interfaces []uint16
	fields []*MemberInfo
	methods []*MemberInfo
	attributes []AttributeInfo
}


// 解析class
func Parser(classData []byte) ( cf *ClassFile, err error) {
	// 可以使发生panic后停止上报 继续执行
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

// 读取解析class
func (c *ClassFile) read(reader *ClassReader){
	c.readAndCheckMagic(reader)
	c.readAndCheckVersion(reader)
	c.constantPool = readConstantPool(reader)
	fmt.Println("cpCount: ", len(c.constantPool), "   ", cap(c.constantPool))
	c.accessFlags = reader.readUint16() // 访问标志
	c.thisClass = reader.readUint16()
	c.superClass = reader.readUint16()
	c.interfaces = reader.readUint16s()
	c.fields = readMembers(reader, c.constantPool)
	c.methods = readMembers(reader, c.constantPool)
	c.attributes = readAttributes(reader, c.constantPool)  // 读取类的所有attribute
}

// 读取并且检查魔数
func (c *ClassFile) readAndCheckMagic(reader *ClassReader){
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {   // 检查魔数是否是cafebabe
		panic("java.lang.classFormatError : magic !")
	}
}

// 读取并且检查版本号
func (c *ClassFile) readAndCheckVersion(reader *ClassReader){
	c.minorVersion = reader.readUint16()
	c.majorVersion = reader.readUint16()
	//支持45-52的版本号 (45有多个小版本号)
	switch c.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if c.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

// getter
func (c *ClassFile) MinorVersion() uint16 {
	return c.minorVersion
}

func (c *ClassFile) MajorVersion() uint16 {
	return c.majorVersion
}

func(c *ClassFile) ClassName() string{
	return c.constantPool.getClassName(c.thisClass)
}

func(c *ClassFile) SuperClassName() string{
	if c.superClass > 0 {
		return c.constantPool.getClassName(c.superClass)
	}
	return "" // 只有Object才没有父类
}

func (c *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(c.interfaces))
	for i, v := range c.interfaces {
		interfaceNames[i] = c.constantPool.getClassName(v)
	}
	return interfaceNames
}

//Getter
func (c *ClassFile) AccessFlags() uint16 {
	return c.accessFlags
}

func (c *ClassFile) ConstantPool() ConstantPool {
	return c.constantPool
}

func (c *ClassFile) Fields() []*MemberInfo {
	return c.fields
}

func (c *ClassFile) Methods() []*MemberInfo {
	return c.methods
}






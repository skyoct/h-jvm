package metaspace

// 判断类之间的关系

// jvms8 6.5.instanceof
// jvms8 6.5.checkcast
func (c *Class) isAssignableFrom(other *Class) bool {
	if c == other {
		return true
	}
	if !other.IsInterface() {
		return c.isSubClassOf(other)
	} else {
		return c.isImplements(other)
	}
}

// 判断c是否继承于某个类
// 递归往上找 找到一个c的一个父类等于other
func (c *Class) isSubClassOf(other *Class) bool {
	for class := c.superClass; class != nil; class = class.superClass {
		if class == other {
			return true
		}
	}
	return false
}

// 判断当前类是否实现一个接口
func (c *Class) isImplements(iface *Class) bool {
	for class := c; class != nil; class = c.superClass { // 循环遍历当前类和父类
		for _, i := range class.interfaces { // 循环遍历接口
			if i == iface || i.isSubInterfaceOf(iface) { // 进行查找
				return true
			}
		}
	}
	return false
}

// 判断当前是否继承自接口
func (c *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range c.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) { // 先查找自己的接口 然后递归查询接口继承的接口
			return true // 接口允许继承多个接口
		}
	}
	return false
}

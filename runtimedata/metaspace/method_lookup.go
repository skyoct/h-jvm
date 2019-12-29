package metaspace

// 在类内查找方法（通过名字和描述符和确定）
func LookupMethodInClass(class *Class, name, descriptor string) *Method { // 在一个类内查找方法
	for c := class; c != nil; c = c.superClass { // 在当前类查找不到去父类查找
		for _, method := range c.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
	}
	return nil
}

// 在接口内查找方法
func lookupMethodInInterfaces(ifaces []*Class, name, descriptor string) *Method {
	for _, iface := range ifaces {
		for _, method := range iface.methods {
			if method.name == name && method.descriptor == descriptor {
				return method
			}
		}
		method := lookupMethodInInterfaces(iface.interfaces, name, descriptor) // 递归向上查找
		if method != nil {
			return method
		}
	}
	return nil
}

package metaspace

type SymRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}

// 如果符号引用已经解析则直接返回 否则解析符号引用
func (s *SymRef) ResolvedClass() *Class {
	if s.class == nil {
		s.resolveClassRef()
	}
	return s.class
}

// 解析符号引用
func (s *SymRef) resolveClassRef() {
	d := s.cp.class
	c := d.classLoader.LoadClass(s.className)
	if !c.isAccessibleTo(d) { // 判断是否能够访问那个类
		panic("java.lang.IllegalAccessError") //类D想访问类C，需要满足两个条件之一：
	} // C是 public，或者C和D在同一个运行时包内

	s.class = c
}

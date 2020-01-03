package metaspace

import "h-jvm/classfile"

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

// 把classfile memberInfo的文件放过来
func (c *ClassMember) copyMemberInfo(info *classfile.MemberInfo) {
	c.accessFlags = info.AccessFlags()
	c.name = info.Name()
	c.descriptor = info.Descriptor()
}

// 计算修饰符 如果按位与 为0  则代表不是
// public的二进制为01  private的二进制为10
// 依次下去  相同的异或保持原值  不同的必然为0
func (c *ClassMember) IsPublic() bool {
	return 0 != c.accessFlags&ACC_PUBLIC
}
func (c *ClassMember) IsPrivate() bool {
	return 0 != c.accessFlags&ACC_PRIVATE
}
func (c *ClassMember) IsProtected() bool {
	return 0 != c.accessFlags&ACC_PROTECTED
}
func (c *ClassMember) IsStatic() bool {
	return 0 != c.accessFlags&ACC_STATIC
}
func (c *ClassMember) IsFinal() bool {
	return 0 != c.accessFlags&ACC_FINAL
}
func (c *ClassMember) IsSynthetic() bool {
	return 0 != c.accessFlags&ACC_SYNTHETIC
}

// getters
func (c *ClassMember) Name() string {
	return c.name
}
func (c *ClassMember) Descriptor() string {
	return c.descriptor
}
func (c *ClassMember) Class() *Class {
	return c.class
}
func (c *Class) Loader() *ClassLoader{
	return c.classLoader
}

// 字段和方法的访问权限判断  判断在d中是否可以访问c的方法
func (c *ClassMember) isAccessibleTo(d *Class) bool {
	if c.IsPublic() {
		return true
	}
	class := c.class
	if c.IsProtected() {
		return d == class || d.IsSubClassOf(class) || class.GetPackageName() == d.GetPackageName()
	}
	if !c.IsPrivate() {
		return class.GetPackageName() == d.GetPackageName()
	}
	return d == class
}

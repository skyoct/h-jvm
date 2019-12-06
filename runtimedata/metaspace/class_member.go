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

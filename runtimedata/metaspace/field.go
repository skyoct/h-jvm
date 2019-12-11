package metaspace

import "h-jvm/classfile"

type Field struct {
	ClassMember
	slotId uint
}

func newField(class *Class, cfFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(cfFields))
	for i, cfField := range cfFields {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyMemberInfo(cfField)
	}
	return fields
}

// long的描述符为J  double的描述符为D  对象的描述符为L
func (f *Field) isDoubleOrLong() bool {
	return f.ClassMember.descriptor == "J" || f.ClassMember.descriptor == "D"
}

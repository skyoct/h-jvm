package metaspace

type Object struct {
	class  *Class
	fields Slots
}

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCount),
	}
}

// getters
func (o *Object) Class() *Class {
	return o.class
}
func (o *Object) Fields() Slots {
	return o.fields
}

// 判断是某个类的实例
func (o *Object) IsInstanceOf(class *Class) bool {
	return class.IsAssignableFrom(o.class)
}

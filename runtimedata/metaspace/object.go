package metaspace

type Object struct {
	class  *Class
	data interface{}
}

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		data: newSlots(class.instanceSlotCount),
	}
}

// getters
func (o *Object) Class() *Class {
	return o.class
}
func (o *Object) Fields() Slots {
	return o.data.(Slots)
}

// 判断是某个类的实例
func (o *Object) IsInstanceOf(class *Class) bool {
	return class.IsAssignableFrom(o.class)
}

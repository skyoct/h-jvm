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

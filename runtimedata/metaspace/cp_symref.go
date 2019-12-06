package metaspace

type SymRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}

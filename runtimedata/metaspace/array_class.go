package metaspace

// 判断是否是数组
func (c *Class) IsArray() bool {
	return c.name[0] == '['
}


func (c *Class) NewArray(count uint) *Object{
	if !c.IsArray() {
		panic("Not array class:" + c.name)
	}
	switch c.Name() {
	case "[Z":
		return &Object{c, make([]int8, count)}
	case "[B":
		return &Object{c, make([]int8, count)}
	case "[C":
		return &Object{c, make([]uint16, count)}
	case "[S":
		return &Object{c, make([]int16, count)}
	case "[I":
		return &Object{c, make([]int32, count)}
	case "[J":
		return &Object{c, make([]int64, count)}
	case "[F":
		return &Object{c, make([]float32, count)}
	case "[D":
		return &Object{c, make([]float64, count)}
	default:
		return &Object{c, make([]*Object, count)}
	}
}

// 加载array class
func (c *Class) ArrayClass() *Class{
	arrayClassName := getArrayClassName(c.name)
	return c.Loader().LoadClass(arrayClassName)
}


// 获取类数组的classname
func getArrayClassName(className string) string {
	return "[" + toDescriptor(className)
}

func toDescriptor(className string) string{
	if className[0] == '['{  // 数组名直接返回
		return className
	}
	if d, ok := primitiveTypes[className]; ok { // 是基本类 返回类型描述符
		return d
	}
	return "L" + className + ";"  // 返回java类型描述符
}

var primitiveTypes = map[string]string{
	"void":    "V",
	"boolean": "Z",
	"byte":    "B",
	"short":   "S",
	"int":     "I",
	"long":    "J",
	"char":    "C",
	"float":   "F",
	"double":  "D",
}

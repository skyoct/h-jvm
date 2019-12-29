package metaspace

import "strings"

type MethodDescriptor struct {
	parameterTypes []string // 参数类型列表
	returnType     string   // 返回值列表
}

func (m *MethodDescriptor) addParameterType(t string) {
	pLen := len(m.parameterTypes)
	if pLen == cap(m.parameterTypes) {
		s := make([]string, pLen, pLen+4)
		copy(s, m.parameterTypes)
		m.parameterTypes = s
	}
	m.parameterTypes = append(m.parameterTypes, t)
}

// 解析描述符为
// 描述符构成可以参考博客
// http://blog.october.ink/2019/12/23/JVM%E4%B9%8B%E7%AC%A6%E5%8F%B7%E5%BC%95%E7%94%A8%E5%92%8C%E7%9B%B4%E6%8E%A5%E5%BC%95%E7%94%A8/
type MethodDescriptorParser struct {
	raw    string
	offset int
	parsed *MethodDescriptor
}

func parseMethodDescriptor(descriptor string) *MethodDescriptor {
	parser := &MethodDescriptorParser{}
	return parser.parse(descriptor)
}

func (m *MethodDescriptorParser) parse(descriptor string) *MethodDescriptor {
	m.raw = descriptor
	m.parsed = &MethodDescriptor{}
	m.startParams()
	m.parseParamTypes()
	m.endParams()
	m.parseReturnType()
	m.finish()
	return m.parsed
}

func (m *MethodDescriptorParser) startParams() {
	if m.readUint8() != '(' {
		m.causePanic()
	}
}
func (m *MethodDescriptorParser) endParams() {
	if m.readUint8() != ')' {
		m.causePanic()
	}
}
func (m *MethodDescriptorParser) finish() {
	if m.offset != len(m.raw) {
		m.causePanic()
	}
}

func (m *MethodDescriptorParser) causePanic() {
	panic("BAD descriptor: " + m.raw)
}

func (m *MethodDescriptorParser) readUint8() uint8 {
	b := m.raw[m.offset]
	m.offset++
	return b
}
func (m *MethodDescriptorParser) unreadUint8() {
	m.offset--
}

func (m *MethodDescriptorParser) parseParamTypes() {
	for {
		t := m.parseFieldType()
		if t != "" {
			m.parsed.addParameterType(t)
		} else {
			break
		}
	}
}

func (m *MethodDescriptorParser) parseReturnType() {
	if m.readUint8() == 'V' {
		m.parsed.returnType = "V"
		return
	}

	m.unreadUint8()
	t := m.parseFieldType()
	if t != "" {
		m.parsed.returnType = t
		return
	}

	m.causePanic()
}

func (m *MethodDescriptorParser) parseFieldType() string {
	switch m.readUint8() {
	case 'B':
		return "B"
	case 'C':
		return "C"
	case 'D':
		return "D"
	case 'F':
		return "F"
	case 'I':
		return "I"
	case 'J':
		return "J"
	case 'S':
		return "S"
	case 'Z':
		return "Z"
	case 'L':
		return m.parseObjectType()
	case '[':
		return m.parseArrayType()
	default:
		m.unreadUint8()
		return ""
	}
}

func (m *MethodDescriptorParser) parseObjectType() string {
	unread := m.raw[m.offset:]
	semicolonIndex := strings.IndexRune(unread, ';') // 找出最后才一个;的位置
	if semicolonIndex == -1 {
		m.causePanic()
		return ""
	} else {
		objStart := m.offset - 1
		objEnd := m.offset + semicolonIndex + 1 // 计算需要截断的长度
		m.offset = objEnd
		descriptor := m.raw[objStart:objEnd]
		return descriptor
	}
}

// 解析数组类型
func (m *MethodDescriptorParser) parseArrayType() string {
	arrStart := m.offset - 1
	m.parseFieldType()
	arrEnd := m.offset
	descriptor := m.raw[arrStart:arrEnd]
	return descriptor
}

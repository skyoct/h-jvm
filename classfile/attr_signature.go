package classfile

/*
Signature_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 signature_index;
}
*/

//jdk1.5 新增的 主要用于支持泛型情况下的签名
type SignatureAttribute struct {
	cp             ConstantPool
	signatureIndex uint16
}

func (s *SignatureAttribute) readInfo(reader *ClassReader) {
	s.signatureIndex = reader.readUint16()
}

func (s *SignatureAttribute) Signature() string {
	return s.cp.getUtf8(s.signatureIndex)
}

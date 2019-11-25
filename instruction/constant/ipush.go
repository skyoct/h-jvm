package constant

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

// 本文件共计2条 （0x10-0x11）

//0x10 bipush    将单字节的常量值(-128~127)推送至栈顶
type BIPush struct {
	val int8
}

func (b *BIPush) FetchOperands(reader *base.CodeReader) {
	b.val = reader.ReadInt8()
}

func (b *BIPush) Execute(frame *runtimedata.Frame) {
	i := int32(b.val)
	frame.OperandStack().PushInt(i)
}

// 0x11 sipush    将一个短整型常量值(-32768~32767)推送至栈顶
type SIPush struct {
	val int16
}

func (s *SIPush) FetchOperands(reader *base.CodeReader) {
	s.val = reader.ReadInt16()
}

func (s *SIPush) Execute(frame *runtimedata.Frame) {
	i := int32(s.val)
	frame.OperandStack().PushInt(i)
}

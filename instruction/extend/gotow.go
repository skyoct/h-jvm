package extend

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

// goto_w    无条件跳转（宽索引）  和goto指令的唯一区别就是索引从2字节变成了4 字节
type GotoW struct {
	offset int
}

func (g *GotoW) FetchOperands(reader *base.CodeReader) {
	g.offset = int(reader.ReadInt32())
}

func (g *GotoW) Execute(frame *runtimedata.Frame) {
	base.Branch(frame, g.offset)
}

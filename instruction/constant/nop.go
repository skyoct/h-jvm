package constant

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

// 本文件共计1条 （0x00）

// 0x00 nop  什么都不做
type Nop struct {
	base.NoOperandsInstruction
}

func (n *Nop) Execute(frame *runtimedata.Frame) {

}

package control

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

type TableSwitch struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (t *TableSwitch) FetchOperands(reader *base.CodeReader) {
	reader.SkipPadding() // 去除padding
	t.defaultOffset = reader.ReadInt32()
	// tableswitch 是连续的 low和high机case对应位置
	t.low = reader.ReadInt32()
	t.high = reader.ReadInt32()
	count := t.high - t.low + 1 // 计算case的个数
	t.jumpOffsets = reader.ReadInt32s(count)
}

func (t *TableSwitch) Execute(frame *runtimedata.Frame) {
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= t.low && index <= t.high {
		offset = int(t.jumpOffsets[index-t.low])
	} else {
		offset = int(t.defaultOffset)
	}
	base.Branch(frame, offset)
}

// 0xab lookupswitch   用于switch条件跳转，case值不连续（可变长度指令）
// lookswitch的结构类似与map
type LookSwitch struct {
	defaultOffset int32
	nPairs        int32   // 代表数量
	matchOffsets  []int32 //存储方式为0为case 1为偏移量 依次类推
}

func (l *LookSwitch) FetchOperands(reader *base.CodeReader) {
	reader.SkipPadding() // 去除padding
	l.nPairs = reader.ReadInt32()
	l.matchOffsets = reader.ReadInt32s(l.nPairs * 2)
}

func (l *LookSwitch) Execute(frame *runtimedata.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < l.nPairs*2; i += 2 {
		if l.matchOffsets[i] == key { // 找到key相同的了
			base.Branch(frame, int(l.matchOffsets[i+1]))
			return
		}
	}
	base.Branch(frame, int(l.defaultOffset)) // 默认分支
}

package math

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

//0x84 iinc     将指定int型变量增加指定值（i++, i--, i+=2）

type IInc struct {
	Index uint
	Const int32
}

func (i *IInc) FetchOperands(reader *base.CodeReader) {
	i.Index = uint(reader.ReadUint8())
	i.Const = int32(reader.ReadInt8())
}

// 取出局部变量表对应位置的值 然后加上指定的值 然后放回局部变量表
func (i *IInc) Execute(frame *runtimedata.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(i.Index)
	val += i.Const
	localVars.SetInt(i.Index, val)
}

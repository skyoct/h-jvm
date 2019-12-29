package constant

import (
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

// 常量推至栈顶 本文件共计15条 （0x01 - 0x0f）

// 0x01 aconst_null 将null推送至栈顶
type AConstNull struct {
	base.NoOperandsInstruction
}

func (a *AConstNull) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushRef(nil)
}

// 0x02 iconst_m1   将int型-1推送至栈顶
type IConstM1 struct {
	base.NoOperandsInstruction
}

func (i *IConstM1) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushInt(-1)
}

// 0x03 iconst_0   将int型0推送至栈顶
type IConst0 struct {
	base.NoOperandsInstruction
}

func (i *IConst0) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushInt(0)
}

// 0x04 iconst_1   将int型1推送至栈顶
type IConst1 struct {
	base.NoOperandsInstruction
}

func (i *IConst1) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushInt(1)
}

// 0x05 iconst_2   将int型2推送至栈顶
type IConst2 struct {
	base.NoOperandsInstruction
}

func (i *IConst2) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushInt(2)
}

// 0x06 iconst_3   将int型3推送至栈顶
type IConst3 struct {
	base.NoOperandsInstruction
}

func (i *IConst3) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushInt(3)
}

// 0x07 iconst_4   将int型4推送至栈顶
type IConst4 struct {
	base.NoOperandsInstruction
}

func (i *IConst4) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushInt(4)
}

// 0x08 iconst_5   将int型5推送至栈顶
type IConst5 struct {
	base.NoOperandsInstruction
}

func (i *IConst5) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushInt(5)
}

//0x09 lconst_0   将long型0推送至栈顶
type LConst0 struct {
	base.NoOperandsInstruction
}

func (l *LConst0) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushLong(0)
}

//0x0a lconst_1   将long型1推送至栈顶
type LConst1 struct {
	base.NoOperandsInstruction
}

func (l *LConst1) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushLong(1)
}

//0x0b fconst_0   将float型0推送至栈顶
type FConst0 struct {
	base.NoOperandsInstruction
}

func (f *FConst0) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

//0x0c fconst_1   将float型1推送至栈顶
type FConst1 struct {
	base.NoOperandsInstruction
}

func (f *FConst1) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

//0x0d fconst_2   将float型2推送至栈顶
type FConst2 struct {
	base.NoOperandsInstruction
}

func (f *FConst2) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

// 0x0e dconst_0   将double型0推送至栈顶
type DConst0 struct {
	base.NoOperandsInstruction
}

func (d *DConst0) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

// 0x0f dconst_1   将double型1推送至栈顶
type DConst1 struct {
	base.NoOperandsInstruction
}

func (d *DConst1) Execute(frame *runtimedata.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

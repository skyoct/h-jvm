package main

import (
	"fmt"
	"h-jvm/classfile"
	"h-jvm/instruction"
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
)

func interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	code := codeAttr.Code()
	thread := runtimedata.NewThread()                  // 创建一个线程
	frame := runtimedata.NewFrame(maxLocals, maxStack) //创建一个帧
	thread.PushFrame(frame)                            // 把帧压入java虚拟机栈中
	defer catchErr(frame)
	loop(thread, code)
}

func catchErr(frame *runtimedata.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *runtimedata.Thread, code []byte) {
	frame := thread.PopFrame()
	reader := &base.CodeReader{}
	for {
		pc := frame.NextPc()
		thread.SetPc(pc)
		reader.Reset(code, pc)
		opcode := reader.ReadUint8()
		inst := instruction.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPc(reader.Pc())
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}

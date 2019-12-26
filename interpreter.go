package main

import (
	"fmt"
	"h-jvm/instruction"
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
	"h-jvm/runtimedata/metaspace"
)

func interpret(method *metaspace.Method) {
	thread := runtimedata.NewThread()
	frame := runtimedata.NewFrame(method)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, method.Code())
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

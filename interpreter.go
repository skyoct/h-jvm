package main

import (
	"fmt"
	"h-jvm/instruction"
	"h-jvm/instruction/base"
	"h-jvm/runtimedata"
	"h-jvm/runtimedata/metaspace"
)

func interpret(method *metaspace.Method, logInst bool) {
	thread := runtimedata.NewThread()
	frame := runtimedata.NewFrame(method)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, logInst)
}

func catchErr(frame *runtimedata.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		logFrames(frame.Thread())
		panic(r)
	}
}

func loop(thread *runtimedata.Thread, logInst bool) {
	reader := &base.CodeReader{}
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPc()
		thread.SetPc(pc)
		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadUint8()
		inst := instruction.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPc(reader.Pc())
		if logInst {
			logInstruction(frame, inst)
		}
		// execute
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}


func logInstruction(frame *runtimedata.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().Pc()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}

func logFrames(thread *runtimedata.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n",
			frame.NextPc(), className, method.Name(), method.Descriptor())
	}
}
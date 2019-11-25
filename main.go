package main

import (
	"fmt"
	"h-jvm/classfile"
	"h-jvm/classpath"
	"h-jvm/runtimedata"
	"strings"
)

func main() {
	cmd := cmdParser()
	if cmd.versionFlag {
		fmt.Print("h-jvm version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {

	cp := classpath.Parser(cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("main method not found")
	}
}

func printStack() {
	thread := runtimedata.NewThread()
	frame := runtimedata.NewFrame(100, 100)
	thread.PushFrame(frame)
	currentFrame := thread.CurrentFrame()
	currentFrame.LocalVars().SetInt(0, 123)
	currentFrame.LocalVars().SetLong(1, 233333333333)
	currentFrame.LocalVars().SetFloat(3, 1.222)
	currentFrame.LocalVars().SetDouble(4, 2.222222222)
	fmt.Println(currentFrame.LocalVars().GetInt(0))
	fmt.Println(currentFrame.LocalVars().GetFloat(3))
	fmt.Println(currentFrame.LocalVars().GetDouble(4))
	fmt.Println(currentFrame.LocalVars().GetLong(1))
	currentFrame.OperandStack().PushInt(123)
	currentFrame.OperandStack().PushFloat(1.2222)
	currentFrame.OperandStack().PushLong(1234567891234)
	currentFrame.OperandStack().PushDouble(1.23456781222)
	fmt.Println(currentFrame.OperandStack().PopDouble())
	fmt.Println(currentFrame.OperandStack().PopLong())
	fmt.Println(currentFrame.OperandStack().PopFloat())
	fmt.Println(currentFrame.OperandStack().PopInt())

}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("class not found or load main class %s", className)
	}
	cf, err := classfile.Parser(classData)
	if err != nil {
		err.Error()
	}
	return cf
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("  %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("  %s\n", m.Name())
		fmt.Println(m.CodeAttribute().Code())

	}
}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range cf.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}

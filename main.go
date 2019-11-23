package main

import (
	"fmt"
	"strings"
)
import "h-jvm/classpath"
import "h-jvm/classfile"

func main(){
	cmd := cmdParser()
	if cmd.versionFlag {
		fmt.Print("h-jvm version 0.0.1")
	}else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd * Cmd) {
	cp := classpath.Parser(cmd.cpOption)
	fmt.Printf("classpath: %v class: %v args:%v \n", cp, cmd.cpOption, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	printClassInfo(cf)
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile{
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
	}
}
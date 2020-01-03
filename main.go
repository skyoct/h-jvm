package main

import (
	"fmt"
	"h-jvm/classpath"
	"h-jvm/runtimedata/metaspace"
)

func main() {
	cmd := cmdParser()
	//cmd.cpOption = "/Users/october"
	//cmd.class = "Fibonacci"
	//cmd.jreOption = "/Users/october/WorkSpace/jre"
	if cmd.helpFlag {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parser(cmd.cpOption, cmd.jreOption)
	classLoader := metaspace.NewClassLoader(cp)
	mainClass := classLoader.LoadClass(cmd.class)
	mainMethod := mainClass.GetMainMethod()
	//fmt.Print(mainClass.Name())
	mainClass.NewObject()
	if mainMethod != nil{
		interpret(mainMethod, false)
	}else{
		fmt.Print("not found")
	}
}

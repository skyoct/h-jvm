package main

import (
	"fmt"
	"h-jvm/classpath"
	"h-jvm/runtimedata/metaspace"
)

func main() {
	startJVM()
}

func startJVM() {
	cp := classpath.Parser("")
	classLoader := metaspace.NewClassLoader(cp)
	mainClass := classLoader.LoadClass("Fibonacci")
	mainMethod := mainClass.GetMainMethod()
	fmt.Print(mainClass.Name())
	mainClass.NewObject()
	if mainMethod != nil{
		interpret(mainMethod, false)
	}else{
		fmt.Print("not found")
	}
}

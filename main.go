package main

import (
	"fmt"
	"h-jvm/classpath"
	"h-jvm/runtimedata/metaspace"
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
	classLoader := metaspace.NewClassLoader(cp)
	className := strings.Replace(cmd.class, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("main method not found")
	}

}

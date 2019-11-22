package main

import (
	"fmt"
	"strings"
)
import "h-jvm/classpath"

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
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("class not found or load main class %s", cmd.class)
	}
	fmt.Printf("class data: %v \n", classData)
}

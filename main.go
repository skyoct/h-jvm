package main

import "fmt"

func main(){
	cmd := cmdParser()
	if cmd.versionFlag {
		fmt.Print("h-jvm version 0.0.1")
	}else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		fmt.Print("start JVM")
	}
}

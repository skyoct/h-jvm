package main

import "flag"
import "fmt"
import "os"

/*
	命令行选项和参数的结构体
 */
type Cmd struct {
	helpFlag bool
	versionFlag bool
	cpOption string
	XjreOption string
	class string
	args []string
}

/*
	解析参数的函数
 */

func cmdParser() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "classpath")

	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func printUsage(){
	fmt.Print("Usage: %s [-options] class [args...]\n", os.Args[0])
}



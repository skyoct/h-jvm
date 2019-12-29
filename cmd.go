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
	jreOption string   //非标注 用来加载jre的类
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
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")  // classpath  和 cp一样
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.jreOption, "jre", "", "path to jre")

	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

// 提示信息
func printUsage(){
	fmt.Printf("Usage: %s [-options] class [args...]\n",  os.Args[0])
}



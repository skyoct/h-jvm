package classpath

import (
	"os"
	"path/filepath"
)

// jre路径
const jrePath = "E:\\Code\\Go\\src\\h-jvm"

type Classpath struct {
	bootClasspath Entry // 启动类加载器
	extClasspath  Entry // 扩展类加载器
	userClasspath Entry // 用户类加载器
}

/**
算是双亲委派
按照启动类加载器 -> 扩展类加载器 -> 用户类加载器 顺序来查找类
*/
func (c *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class" // 添加以class结尾
	if data, entry, err := c.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := c.userClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return c.userClasspath.readClass(className)
}

func (c *Classpath) String() string {
	return c.userClasspath.String()
}

func Parser(cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath() // 解析启动类和扩展类的加载器
	cp.parserUserClasspath(cpOption)
	return cp
}

/**
解析启动类和扩展类
*/

func (c *Classpath) parseBootAndExtClasspath() {
	if !exists(jrePath) {
		panic("jre不存在......")
	}
	// jre/lib/*
	jreLibPath := filepath.Join(jrePath, "lib", "*")
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jrePath, "lib", "ext", "*")
	c.bootClasspath = newWildcardEntry(jreLibPath)
	c.extClasspath = newWildcardEntry(jreExtPath)
}

/**
解析用户类
*/

func (c *Classpath) parserUserClasspath(cpOption string) {

	// 如果用户没指定cp 则使用当前目录作为用户类的路径
	if cpOption == "" {
		cpOption = "."
	}
	c.userClasspath = newEntry(cpOption)
}

/*
	判断文件或者文件夹是否存在
*/
func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

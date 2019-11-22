package classpath

import (
	"os"
	"strings"
)

// 文件分隔符 在读取类时候来分割通过文件分隔符分割的多个文件
const pathListSeparator  = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)     //参数是类的相对路径,以.class 结尾 比如java/lang/Object.class
	String() string   // 返回变量以字符串表示
}

func newEntry(path string) Entry{
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	// 以jar 和 zip结尾的
	if strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") ||
		strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
		newZipEntry(path)
	}
	return newDirEntry(path)
}

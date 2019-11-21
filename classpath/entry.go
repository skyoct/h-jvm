package classpath

import (
	"os"
)

// 文件分隔符 在读取类时候来分割通过文件分隔符分割的多个文件
const pathListSeparator  = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry{
	return newDirEntry(path)
}

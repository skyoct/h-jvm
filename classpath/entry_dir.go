package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absPath string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

func(d * DirEntry) readClass(className string) ([]byte, Entry, error){
	fileName := filepath.Join(d.absPath, className)
	data, err := ioutil.ReadFile(fileName)
	return data, d, err
}

func (d DirEntry) String() string {
	return d.absPath
}
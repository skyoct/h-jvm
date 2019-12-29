package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry


func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	// 通过分隔符分割
	for _, path := range strings.Split(pathList, pathListSeparator){
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (z CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range z {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found:" + className)
}

func (z CompositeEntry) String() string {
	strs := make([]string, len(z))
	for i, entry := range z {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}




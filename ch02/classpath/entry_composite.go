package classpath

import (
	"errors"
	"strings"
)

// 多个Entry组成, 它只是寻找Entry, 不进行数据的读取
type CompositeEntry []Entry

// 根据分隔符去分开Entry
func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (this CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range this {
		date, from, err := entry.readClass(className)
		if err == nil {
			return date, from, nil
		}
	}
	return nil, nil, errors.New("composite class not found: " + className)
}

func (this CompositeEntry) String() string {
	strs := make([]string, len(this))
	for i, entry := range this {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}

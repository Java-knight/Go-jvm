package classpath

import (
	"os"
	"strings"
)

// 加载类路径的分隔符（mac中是"/", win是"\"）
const pathListSeparator = string(os.PathSeparator)

type Entry interface {
	// 寻找 和 加载class文件
	readClass(className string) ([]byte, Entry, error)
	// 类似Java中的toString()
	String() string
}

// 根据参数创建不同类型的Entry实例
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}

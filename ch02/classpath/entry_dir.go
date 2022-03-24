package classpath

import (
	"io/ioutil"
	"path/filepath"
)

// 目录形式的类路径
type DirEntry struct {
	absDir string
}

// 创建DirEntry实例（相当于构造函数）
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path) // 将path转化为 绝对路径
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

// 读取类文件的内容（通过path+className找到类, ioutil读取）
func (this *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(this.absDir, className)
	date, err := ioutil.ReadFile(fileName)
	return date, this, err
}

// toString()方法, 返回path的 绝对路径
func (this *DirEntry) String() string {
	return this.absDir
}

package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

// 处理以 ".jar\.JAR\.zip\.ZIP" 结尾的路径
type ZipEntry struct {
	absPath string
}

// 创建ZipEntry实例
func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

/*
 读取压缩文件的class文件的内容
*/
func (this *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	reader, err := zip.OpenReader(this.absPath)
	if err != nil {
		return nil, nil, err
	}

	defer reader.Close()
	// 读取压缩包里面的文件和className名字相同的文件（xxx.class, xxx就是className）
	for _, file := range reader.File {
		if file.Name == className {
			classReader, err := file.Open()
			if err != nil {
				return nil, nil, err
			}

			defer classReader.Close()
			date, err := ioutil.ReadAll(classReader)
			if err != nil {
				return nil, nil, err
			}
			return date, this, nil

		}
	}
	return nil, nil, errors.New("zip class not found: " + className)
}

func (this *ZipEntry) String() string {
	return this.absPath
}

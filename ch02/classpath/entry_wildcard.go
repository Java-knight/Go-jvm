package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

// WildcardEntry其实就是CompositeEntry的一种
// 删除path后面的*, 调用Walk函数, 扫描path路径下的.jar\.JAR, 遇到文件或者
func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1] // remove last "*"
	compositeEntry := []Entry{}

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 当找到.jar文件就不扫描下面的目录, 直接跳出
		if info.IsDir() && path != baseDir { // info.IsDir()=true就是路径, 否则就是父路径
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	filepath.Walk(baseDir, walkFn) // 以baseDir根的一个文件目录树, 每个目录都需要访问walkFn
	return compositeEntry
}

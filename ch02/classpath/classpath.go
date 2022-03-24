package classpath

import (
	"os"
	"path/filepath"
)

// 类路径
type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

/*
 解析函数,
    (1) -Xjre解析启动类路径和扩展类路径
    (2) -classpath/-cp 解析用户类路径
*/
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

// 如果没有输入-classpath/-cp, 就使用当前目录作为用户类路径（HotSpot VM也是）, 依次启动类路径、扩展类路径、用户类路径
// 注意，传递给ReadClass函数的类名不包含".class"后缀, 这里需要加一下
func (this *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if date, entry, err := this.bootClasspath.readClass(className); err == nil {
		return date, entry, err
	}
	if date, entry, err := this.extClasspath.readClass(className); err == nil {
		return date, entry, err
	}
	return this.userClasspath.readClass(className)

}

// toString方法返回用户类路径
func (this *Classpath) String() string {
	return this.userClasspath.String()
}

// 解析启动类路径 和 扩展类路径, 使用WildcardEntry解析
func (this *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	// jre/bin/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	this.bootClasspath = newWildcardEntry(jreLibPath)

	// jre/bin/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	this.extClasspath = newWildcardEntry(jreExtPath)
}

// 解析用户类路径
func (this *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	this.userClasspath = newEntry(cpOption)
}

// 优先使用-Xjre作为jre目录; 如果没有输入该选项, 则再当前目录寻找jre目录; 如果找不到, 尝试使用JAVA_HOME环境变量
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

// 判断目录是否存在
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

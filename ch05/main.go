package main

import (
	"Go-jvm/ch05/classfile"
	"Go-jvm/ch05/classpath"
	"fmt"
	"strings"
)

// 启动JVM
func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classFile := loadClass(className, cp)
	mainMethod := getMainMethod(classFile)
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}

// 加载class文件(将内存中的[]byte转换成classFile)
func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	classFile, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return classFile
}

// 获取main方法
func getMainMethod(classFile *classfile.ClassFile) *classfile.MemberInfo {
	for _, method := range classFile.Methods() {
		if method.Name() == "main" && method.Descriptor() == "([Ljava/lang/String;)V" {
			return method
		}
	}
	return nil
}

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

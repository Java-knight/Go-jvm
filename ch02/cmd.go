package main

import (
	"flag"
	"fmt"
	"os"
)

/*
Java的命令行
 -version 输出版本信息, 然后退出
 -?/help: 输出帮助信息, 然后退出
 -cp/-classpath: 指定用户类路径
 -Dproperty=value: 设置Java系统属性
 -Xms<size>: 设置初始堆大小
 -Xmx<size>: 设置堆的最大值
 -Xss<size>: 设置虚拟机栈大小
*/
type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	class       string
	args        []string
	XjreOption  string // 指定类的位置
}

// 解析cmd
func parseCmd() *Cmd {
	cmd := &Cmd{}

	// 将cmd输入的值赋值给Usage进行校验
	flag.Usage = printUsage
	// 校验命令行参数
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse() // 校验

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

// 打印cmd的输入值, 提示用户输入格式
func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}

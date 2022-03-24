package main

import (
	"fmt"
	"strings"

	"Go-jvm/ch02/classpath"
)

// 启动JVM
func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v agrs:%v\n", cp, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	classDate, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s, error: %s\n", cmd.class, err)
		return
	}
	fmt.Printf("class date:%v\n", classDate)

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

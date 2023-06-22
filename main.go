package main

import (
	"fmt"
	"strings"
)
import "nilknow-jvm/classpath"
import "nilknow-jvm/cmd"

func main() {
	command := cmd.ParseCmd()
	if command.VersionFlag {
		fmt.Println("version 0.0.1")
	} else if command.HelpFlag || command.Class == "" {
		cmd.PrintUsage()
	} else {
		startJVM(command)
	}
}

func startJVM(cmd *cmd.Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.CpOption)
	fmt.Printf("classpath: %v\nclass: %v\nargs: %v\n", cp, cmd.Class, cmd.Args)

	className := strings.Replace(cmd.Class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.Class)
		return
	}

	fmt.Printf("class data:%v\n", classData)
}

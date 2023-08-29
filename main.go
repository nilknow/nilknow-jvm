package main

import (
	"fmt"
	"nilknow-jvm/classfile"
	"nilknow-jvm/classpath"
	"nilknow-jvm/cmd"
	"strings"
)

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
	fmt.Printf("class: %v\nargs: %v\n", cmd.Class, cmd.Args)

	className := strings.Replace(cmd.Class, ".", "/", -1)
	cf := loadClass(className, cp)
	mainMethod := getMainMethod(cf)
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.Class)
	}
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", className)
		panic(err)
	}

	fmt.Printf("class data:%v\n", classData)
	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("class magic number:%x\n", cf.Magic())
	fmt.Printf("minor version:%d\n", cf.MinorVersion())
	fmt.Printf("major version:%d\n", cf.MajorVersion())
	return cf
}

func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range cf.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}

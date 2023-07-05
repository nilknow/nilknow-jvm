package main

import (
	"nilknow-jvm/cmd"
	"testing"
)

func TestStartJvm(t *testing.T) {
	startJVM(&cmd.Cmd{CpOption: "./java",Class: "HelloWorld"})
}
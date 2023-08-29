package math

import (
	"nilknow-jvm/instruction/interface"
	"nilknow-jvm/rt"
)

type Iand struct {
	base.NoOperandsInstruction
}

type Land struct {
	base.NoOperandsInstruction
}

func (iand *Iand) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	result := val1 & val2
	stack.PushInt(result)
}

func (land *Land) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopLong()
	val1 := stack.PopLong()
	result := val1 & val2
	stack.PushLong(result)
}

package stack

import (
	"nilknow-jvm/instruction/interface"
	"nilknow-jvm/rt"
)

type Pop struct {
	base.NoOperandsInstruction
}

type Pop2 struct {
	base.NoOperandsInstruction
}

func (p *Pop) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

func (p *Pop2) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}

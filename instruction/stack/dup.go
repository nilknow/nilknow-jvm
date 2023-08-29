package stack

import (
	"nilknow-jvm/instruction/interface"
	"nilknow-jvm/rt"
)

type Dup struct {
	base.NoOperandsInstruction
}

type Dup_x1 struct {
	base.NoOperandsInstruction
}

type Dup_x2 struct {
	base.NoOperandsInstruction
}

type Dup2 struct {
	base.NoOperandsInstruction
}

type Dup2_x1 struct {
	base.NoOperandsInstruction
}

type Dup2_x2 struct {
	base.NoOperandsInstruction
}

func (d *Dup) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}



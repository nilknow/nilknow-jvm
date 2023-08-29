package stack

import (
	"nilknow-jvm/instruction/interface"
	"nilknow-jvm/rt"
)

type Swap struct {
	base.NoOperandsInstruction
}

func (s *Swap) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}

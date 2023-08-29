package comparison

import (
	"nilknow-jvm/instruction/interface"
	"nilknow-jvm/rt"
)

type Icmp struct {
	base.NoOperandsInstruction
}

func (icmp *Icmp) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 > val2 {
		stack.PushInt(1)
	} else if val1 == val2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}

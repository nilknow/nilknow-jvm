package comparison

import (
	"nilknow-jvm/instruction/interface"
	"nilknow-jvm/rt"
)

type LCMP struct {
	base.NoOperandsInstruction
}

func (lcmp *LCMP) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopLong()
	val1 := stack.PopLong()
	if val1 > val2 {
		stack.PushInt(1)
	} else if val1 == val2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}

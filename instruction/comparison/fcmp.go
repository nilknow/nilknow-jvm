package comparison

import (
	"nilknow-jvm/instruction/interface"
	"nilknow-jvm/rt"
)

type FCMPG struct {
	base.NoOperandsInstruction
}

type FCMPL struct {
	base.NoOperandsInstruction
}

func (fcmpg *FCMPG) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopFloat()
	val1 := stack.PopFloat()
	if val1 > val2 {
		stack.PushInt(1)
	} else if val1 == val2 {
		stack.PushInt(0)
	} else if val1 < val2 {
		stack.PushInt(-1)
	} else {
		stack.PushInt(1)
	}
}

func (fcmpl *FCMPL) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopFloat()
	val1 := stack.PopFloat()
	if val1 > val2 {
		stack.PushInt(1)
	} else if val1 == val2 {
		stack.PushInt(0)
	} else if val1 < val2 {
		stack.PushInt(-1)
	} else {
		stack.PushInt(-1)
	}
}

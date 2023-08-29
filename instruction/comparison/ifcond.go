package comparison

import (
	"nilknow-jvm/instruction/base"
	"nilknow-jvm/rt"
)

type IFEQ struct {
	base.BranchInstruction
}

type IFNE struct {
	base.BranchInstruction
}

type IFLT struct {
	base.BranchInstruction
}

type IFLE struct {
	base.BranchInstruction
}

type IFGT struct {
	base.BranchInstruction
}

type IFGE struct {
	base.BranchInstruction
}

func (eq *IFEQ) Execute(frame *rt.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, eq.Offset)
	}
}

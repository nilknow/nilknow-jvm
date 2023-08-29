package comparison

import (
	"nilknow-jvm/instruction/base"
	"nilknow-jvm/rt"
)

type IF_ACMPEQ struct {
	base.BranchInstruction
}

type IF_ACMPNE struct {
	base.BranchInstruction
}

func (eq *IF_ACMPEQ) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 == ref2 {
		base.Branch(frame, eq.Offset)
	}
}

func (ne *IF_ACMPNE) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 == ref2 {
		base.Branch(frame, ne.Offset)
	}
}

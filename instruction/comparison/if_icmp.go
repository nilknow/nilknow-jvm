package comparison

import (
	"nilknow-jvm/instruction/base"
	"nilknow-jvm/rt"
)

type IF_ICMPEQ struct {
	base.BranchInstruction
}

type IF_ICMPNE struct {
	base.BranchInstruction
}

type IF_ICMPLT struct {
	base.BranchInstruction
}

type IF_ICMPLE struct {
	base.BranchInstruction
}

type IF_ICMPGT struct {
	base.BranchInstruction
}

type IF_ICMPGE struct {
	base.BranchInstruction
}

func (eq *IF_ICMPEQ) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 == val2 {
		base.Branch(frame, eq.Offset)
	}
}

func (ne *IF_ICMPNE) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 != val2 {
		base.Branch(frame, ne.Offset)
	}
}

func (lt *IF_ICMPLT) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 < val2 {
		base.Branch(frame, lt.Offset)
	}
}

func (le *IF_ICMPLE) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 <= val2 {
		base.Branch(frame, le.Offset)
	}
}

func (gt *IF_ICMPGT) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 > val2 {
		base.Branch(frame, gt.Offset)
	}
}

func (ge *IF_ICMPGE) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 >= val2 {
		base.Branch(frame, ge.Offset)
	}
}
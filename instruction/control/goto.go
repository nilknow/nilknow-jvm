package control

import (
	"nilknow-jvm/instruction/base"
	"nilknow-jvm/rt"
)

type GOTO struct {
	base.BranchInstruction
}

func (g *GOTO) Execute(frame *rt.Frame) {
	base.Branch(frame,g.Offset)
}

package extend

import (
	"nilknow-jvm/instruction/base"
	"nilknow-jvm/rt"
)

type IFNULL struct{ base.BranchInstruction } // Branch if reference is null
type IFNONNULL struct{ base.BranchInstruction }

func (i *IFNULL) Execute(frame rt.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, i.Offset)
	}
}

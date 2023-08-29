package constant

import (
	"nilknow-jvm/instruction/base"
	"nilknow-jvm/rt"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (nop *NOP) Execute(frame *rt.Frame)  {
	//do nothing
}

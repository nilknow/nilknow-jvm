package instruction

import (
	"fmt"
	"nilknow-jvm/instruction/base"
	"nilknow-jvm/instruction/constant"
)

var (
	nop        = &constant.NOP{}
	aConstNull = &constant.ACONST_NULL{}
)

func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00:
		return nop
	//case 0x01:
	//	return aConstNull
	default:
		panic(fmt.Errorf("Unsupported opcode: 0x%x!", opcode))
	}
}


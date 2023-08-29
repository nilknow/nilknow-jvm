package load

import (
	"nilknow-jvm/instruction/interface"
	"nilknow-jvm/rt"
)

type FLOAD struct {
	base.Index8Instruction
}

type FLOAD_0 struct {
	base.NoOperandsInstruction
}

type FLOAD_1 struct {
	base.NoOperandsInstruction
}

type FLOAD_2 struct {
	base.NoOperandsInstruction
}

type FLOAD_3 struct {
	base.NoOperandsInstruction
}

func _fload(frame rt.Frame, index uint) {
	val:=frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

func (fl *FLOAD) Execute(frame rt.Frame) {
	_fload(frame, fl.Index)
}

func (fl *FLOAD_0) Execute(frame rt.Frame) {
	_fload(frame,0)
}

func (fl *FLOAD_1) Execute(frame rt.Frame) {
	_fload(frame,1)
}

func (fl *FLOAD_2) Execute(frame rt.Frame) {
	_fload(frame,2)
}

func (fl *FLOAD_3) Execute(frame rt.Frame) {
	_fload(frame,3)
}

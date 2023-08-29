package load

import (
	"nilknow-jvm/instruction/interface"
	"nilknow-jvm/rt"
)

type DLOAD struct {
	base.Index8Instruction
}

type DLOAD_0 struct {
	base.NoOperandsInstruction
}

type DLOAD_1 struct {
	base.NoOperandsInstruction
}

type DLOAD_2 struct {
	base.NoOperandsInstruction
}

type DLOAD_3 struct {
	base.NoOperandsInstruction
}

func _dload(frame rt.Frame, index uint) {
	val:=frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

func (dl *DLOAD) Execute(frame rt.Frame) {
	_dload(frame, dl.Index)
}

func (dl *DLOAD_0) Execute(frame rt.Frame) {
	_dload(frame,0)
}

func (dl *DLOAD_1) Execute(frame rt.Frame) {
	_dload(frame,1)
}

func (dl *DLOAD_2) Execute(frame rt.Frame) {
	_dload(frame,2)
}

func (dl *DLOAD_3) Execute(frame rt.Frame) {
	_dload(frame,3)
}

package load

import (
	"nilknow-jvm/instruction/interface"
	"nilknow-jvm/rt"
)

type LLOAD struct {
	base.Index8Instruction
}

type LLOAD_0 struct {
	base.NoOperandsInstruction
}

type LLOAD_1 struct {
	base.NoOperandsInstruction
}

type LLOAD_2 struct {
	base.NoOperandsInstruction
}

type LLOAD_3 struct {
	base.NoOperandsInstruction
}

func _lload(frame rt.Frame, index uint) {
	val:=frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

func (ll *LLOAD) Execute(frame rt.Frame) {
	_lload(frame, ll.Index)
}

func (ll *LLOAD_0) Execute(frame rt.Frame) {
	_lload(frame,0)
}

func (ll *LLOAD_1) Execute(frame rt.Frame) {
	_lload(frame,1)
}

func (ll *LLOAD_2) Execute(frame rt.Frame) {
	_lload(frame,2)
}

func (ll *LLOAD_3) Execute(frame rt.Frame) {
	_lload(frame,3)
}

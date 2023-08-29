package load

import (
	"nilknow-jvm/instruction/interface"
	"nilknow-jvm/rt"
)

type ILOAD struct {
	base.Index8Instruction
}

type ILOAD_0 struct {
	base.NoOperandsInstruction
}

type ILOAD_1 struct {
	base.NoOperandsInstruction
}

type ILOAD_2 struct {
	base.NoOperandsInstruction
}

type ILOAD_3 struct {
	base.NoOperandsInstruction
}

func _iload(frame rt.Frame, index uint) {
	val:=frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

func (il *ILOAD) Execute(frame rt.Frame) {
	_iload(frame, il.Index)
}

func (il *ILOAD_0) Execute(frame rt.Frame) {
	_iload(frame,0)
}

func (il *ILOAD_1) Execute(frame rt.Frame) {
	_iload(frame,1)
}

func (il *ILOAD_2) Execute(frame rt.Frame) {
	_iload(frame,2)
}

func (il *ILOAD_3) Execute(frame rt.Frame) {
	_iload(frame,3)
}

package store

import (
	"nilknow-jvm/instruction/interface"
	"nilknow-jvm/rt"
)

type LSTORE struct {
	base.Index8Instruction
}

type LSTORE_0 struct {
	base.NoOperandsInstruction
}

type LSTORE_1 struct {
	base.NoOperandsInstruction
}

type LSTORE_2 struct {
	base.NoOperandsInstruction
}

type LSTORE_3 struct {
	base.NoOperandsInstruction
}

func _lstroe(frame rt.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

func (ls *LSTORE) Execute(frame rt.Frame) {
	_lstroe(frame, uint(ls.Index))
}

func (ls *LSTORE_0) Execute(frame rt.Frame) {
	_lstroe(frame, 0)
}

func (ls *LSTORE_1) Execute(frame rt.Frame) {
	_lstroe(frame, 1)
}

func (ls *LSTORE_2) Execute(frame rt.Frame) {
	_lstroe(frame, 2)
}

func (ls *LSTORE_3) Execute(frame rt.Frame) {
	_lstroe(frame, 3)
}

package store

import (
	"nilknow-jvm/instruction/interface"
	"nilknow-jvm/rt"
)

type ISTORE struct {
	base.Index8Instruction
}

type ISTORE_0 struct {
	base.NoOperandsInstruction
}

type ISTORE_1 struct {
	base.NoOperandsInstruction
}

type ISTORE_2 struct {
	base.NoOperandsInstruction
}

type ISTORE_3 struct {
	base.NoOperandsInstruction
}

func _istroe(frame rt.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

func (is *ISTORE) Execute(frame rt.Frame) {
	_istroe(frame, uint(is.Index))
}

func (is *ISTORE_0) Execute(frame rt.Frame) {
	_istroe(frame, 0)
}

func (is *ISTORE_1) Execute(frame rt.Frame) {
	_istroe(frame, 1)
}

func (is *ISTORE_2) Execute(frame rt.Frame) {
	_istroe(frame, 2)
}

func (is *ISTORE_3) Execute(frame rt.Frame) {
	_istroe(frame, 3)
}

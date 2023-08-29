package conversion

import (
	"nilknow-jvm/instruction/interface"
	"nilknow-jvm/rt"
)

type D2i struct {
	base.NoOperandsInstruction
}

type D2l struct {
	base.NoOperandsInstruction
}

type D2f struct {
	base.NoOperandsInstruction
}

func (d2i *D2i) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

func (d2l *D2l) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int64(d)
	stack.PushLong(i)
}

func (d2f *D2f) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := float32(d)
	stack.PushFloat(i)
}

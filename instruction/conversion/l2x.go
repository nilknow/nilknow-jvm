package conversion

import (
	"nilknow-jvm/instruction/interface"
	"nilknow-jvm/rt"
)

type L2i struct {
	base.NoOperandsInstruction
}

type L2f struct {
	base.NoOperandsInstruction
}

type L2d struct {
	base.NoOperandsInstruction
}

func (l2i *L2i) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	d := stack.PopLong()
	i := int32(d)
	stack.PushInt(i)
}
func (l2f *L2f) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	d := stack.PopLong()
	i := float32(d)
	stack.PushFloat(i)
}

func (l2d *L2d) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	d := stack.PopLong()
	i := float64(d)
	stack.PushDouble(i)
}

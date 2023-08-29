package conversion

import (
	"nilknow-jvm/instruction/interface"
	"nilknow-jvm/rt"
)

type F2i struct {
	base.NoOperandsInstruction
}

type F2l struct {
	base.NoOperandsInstruction
}

type F2d struct {
	base.NoOperandsInstruction
}

func (f2i *F2i) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	d := stack.PopFloat()
	i := int32(d)
	stack.PushInt(i)
}

func (f2l *F2l) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	d := stack.PopFloat()
	i := int64(d)
	stack.PushLong(i)
}

func (f2d *F2d) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	d := stack.PopFloat()
	i := float64(d)
	stack.PushDouble(i)
}

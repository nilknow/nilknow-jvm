package conversion

import (
	"nilknow-jvm/instruction/interface"
	"nilknow-jvm/rt"
)

type I2l struct {
	base.NoOperandsInstruction
}

type I2f struct {
	base.NoOperandsInstruction
}

type I2d struct {
	base.NoOperandsInstruction
}

func (i2l *I2l) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	d := stack.PopInt()
	i := int64(d)
	stack.PushLong(i)
}

func (i2f *I2f) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	d := stack.PopInt()
	i := float32(d)
	stack.PushFloat(i)
}

func (i2d *I2d) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	d := stack.PopInt()
	i := float64(d)
	stack.PushDouble(i)
}

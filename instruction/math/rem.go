package math

import (
	"math"
	"nilknow-jvm/instruction/interface"
	"nilknow-jvm/rt"
)

type Drem struct {
	base.NoOperandsInstruction
}

type Frem struct {
	base.NoOperandsInstruction
}

type Irem struct {
	base.NoOperandsInstruction
}

type Lrem struct {
	base.NoOperandsInstruction
}

func (ir *Irem) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := val1 % val2
	stack.PushInt(result)
}

func (lr *Lrem) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopLong()
	val1 := stack.PopLong()
	if val2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := val1 / val2
	stack.PushLong(result)
}

func (fr *Frem) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopFloat()
	val1 := stack.PopFloat()
	result := math.Mod(float64(val1), float64(val2))
	stack.PushFloat(float32(result))
}

func (dr *Drem) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopDouble()
	val1 := stack.PopDouble()
	result := math.Mod(val1, val2)
	stack.PushDouble(result)
}

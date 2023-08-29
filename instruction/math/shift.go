package math

import (
	"nilknow-jvm/instruction/interface"
	"nilknow-jvm/rt"
)

type Ishl struct {
	base.NoOperandsInstruction
}

type Ishr struct {
	base.NoOperandsInstruction
}

type Iushr struct {
	base.NoOperandsInstruction
}

type Lshl struct {
	base.NoOperandsInstruction
}

type Lshr struct {
	base.NoOperandsInstruction
}

type Lushr struct {
	base.NoOperandsInstruction
}

func (ishl *Ishl) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	s := uint32(val2) & 0x1f //5bit
	result := val1 >> s
	stack.PushInt(result)
}

func (ishr *Ishr) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	s := uint32(val2) & 0x1f //5bit
	result := val1 >> s
	stack.PushInt(result)
}

func (iushr *Iushr) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	s := uint32(val2) & 0x1f //5bit
	result := int32(uint32(val1)) >> s
	stack.PushInt(result)
}

func (lshl *Lshl) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopLong()
	s := uint32(val2) & 0x3f //6bit
	result := val1 << s
	stack.PushLong(result)
}

func (lshr *Lshr) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopLong()
	s := uint32(val2) & 0x3f //6bit
	result := val1 >> s
	stack.PushLong(result)
}

func (lushr *Lushr) Execute(frame rt.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopLong()
	s := uint32(val2) & 0x3f //6bit
	result := int64(uint64(val1)) >> s
	stack.PushLong(result)
}

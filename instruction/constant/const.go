package constant

import (
	"nilknow-jvm/instruction/base"
	"nilknow-jvm/rt"
)

type ACONST_NULL struct {
	base.NoOperandsInstruction
}

type DConst0 struct {
	base.NoOperandsInstruction
}

type DConst1 struct {
	base.NoOperandsInstruction
}

type FConst0 struct {
	base.NoOperandsInstruction
}

type FConst1 struct {
	base.NoOperandsInstruction
}

type FConst2 struct {
	base.NoOperandsInstruction
}

type IConst0 struct {
	base.NoOperandsInstruction
}

type IConstM1 struct {
	base.NoOperandsInstruction
}

type IConst1 struct {
	base.NoOperandsInstruction
}

type IConst2 struct {
	base.NoOperandsInstruction
}

type IConst3 struct {
	base.NoOperandsInstruction
}

type IConst4 struct {
	base.NoOperandsInstruction
}

type IConst5 struct {
	base.NoOperandsInstruction
}

type LConst0 struct {
	base.NoOperandsInstruction
}

type LConst1 struct {
	base.NoOperandsInstruction
}

func (acn *ACONST_NULL) Execute(frame rt.Frame) {
	frame.OperandStack().PushRef(nil)
}

func (dc0 *DConst0) Execute(frame rt.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

func (acn *DConst1) Execute(frame rt.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

func (acn *FConst0) Execute(frame rt.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

func (acn *FConst1) Execute(frame rt.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

func (acn *FConst2) Execute(frame rt.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

func (acn *IConstM1) Execute(frame rt.Frame) {
	frame.OperandStack().PushInt(-1)
}

func (acn *IConst0) Execute(frame rt.Frame) {
	frame.OperandStack().PushInt(0)
}

func (acn *IConst1) Execute(frame rt.Frame) {
	frame.OperandStack().PushInt(1)
}

func (acn *IConst2) Execute(frame rt.Frame) {
	frame.OperandStack().PushInt(2)
}

func (acn *IConst3) Execute(frame rt.Frame) {
	frame.OperandStack().PushInt(3)
}

func (acn *IConst4) Execute(frame rt.Frame) {
	frame.OperandStack().PushInt(4)
}

func (acn *IConst5) Execute(frame rt.Frame) {
	frame.OperandStack().PushInt(5)
}

func (acn *LConst0) Execute(frame rt.Frame) {
	frame.OperandStack().PushLong(0)
}

func (acn *LConst1) Execute(frame rt.Frame) {
	frame.OperandStack().PushLong(1)
}



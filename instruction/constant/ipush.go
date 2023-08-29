package constant

import (
	"nilknow-jvm/instruction/base"
	"nilknow-jvm/rt"
)

type Bipush struct {
	val int8
}

type Sipush struct {
	val int16
}

func (bp *Bipush) FetchOperands(reader *base.BytecodeReader) {
	bp.val=reader.ReadInt8()
}

func (bp *Bipush) Execute(frame *rt.Frame) {
	i:=int32(bp.val)
	frame.OperandStack().PushInt(i)
}

func (sp *Sipush) FetchOperands(reader *base.BytecodeReader) {
	sp.val=reader.ReadInt16()
}

func (sp *Sipush) Execute(frame *rt.Frame) {
	i:=int32(sp.val)
	frame.OperandStack().PushInt(i)
}

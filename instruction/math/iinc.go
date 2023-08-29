package math

import (
	"nilknow-jvm/instruction/base"
	"nilknow-jvm/rt"
)

type IINC struct {
	Index uint
	Const int32
}

func (iinc *IINC) FetchOperands(reader base.BytecodeReader) {
	iinc.Index = uint(reader.ReadUint8())
	iinc.Const = int32(reader.ReadUint8())
}

func (iinc *IINC) Execute(frame rt.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(iinc.Index)
	val += iinc.Const
	localVars.SetInt(iinc.Index, val)
}

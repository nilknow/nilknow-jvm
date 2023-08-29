package base

import (
	"nilknow-jvm/rt"
)

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rt.Frame)
}

type NoOperandsInstruction struct{}

func (noi *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {

}

type BranchInstruction struct {
	Offset int
}

func (b *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	b.Offset = int(reader.ReadInt16())
}

type JumpInstruction struct {
	Offset int
}

func (ji *JumpInstruction) FetchOperands(reader *BytecodeReader) {
	ji.Offset = int(reader.ReadInt16())
}

type Index8Instruction struct {
	Index uint
}

func (i8i *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	i8i.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint
}

func (i16i *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	i16i.Index = uint(reader.ReadUint16())
}

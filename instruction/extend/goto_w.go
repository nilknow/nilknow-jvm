package extend

import (
	"nilknow-jvm/instruction/base"
	"nilknow-jvm/rt"
)

type GOTO_W struct {
	offset int
}

func (g *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	g.offset = int(reader.ReadInt32())
}
func (g *GOTO_W) Execute(frame *rt.Frame) {
	base.Branch(frame, g.offset)
}

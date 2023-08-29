package control

import (
	"nilknow-jvm/instruction/base"
	"nilknow-jvm/rt"
)

type LOOKUP_SWITCH struct {
	defaultOffset int32
	nPairs        int32
	matchOffsets  []int32
}

func (l *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	l.defaultOffset = reader.ReadInt32()
	l.nPairs = reader.ReadInt32()
	l.matchOffsets = reader.ReadInt32s(l.nPairs * 2)
}

func (l *LOOKUP_SWITCH) Execute(frame *rt.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < l.nPairs*2; i += 2 {
		if l.matchOffsets[i] == key {
			offset := l.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(l.defaultOffset))
}

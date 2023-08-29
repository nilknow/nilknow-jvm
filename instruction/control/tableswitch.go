package control

import (
	"nilknow-jvm/instruction/base"
	"nilknow-jvm/rt"
)

type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

func (t *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	t.defaultOffset = reader.ReadInt32()
	t.low = reader.ReadInt32()
	t.high = reader.ReadInt32()
	jumpOffsetsCount := t.high - t.low + 1
	t.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (self *TABLE_SWITCH) Execute(frame *rt.Frame) {
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffsets[index-self.low])
	} else {
		offset = int(self.defaultOffset)
	}
	base.Branch(frame, offset)
}

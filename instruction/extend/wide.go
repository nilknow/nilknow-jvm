package extend

import (
	"nilknow-jvm/instruction/base"
	"nilknow-jvm/instruction/load"
	"nilknow-jvm/instruction/math"
	"nilknow-jvm/instruction/store"
	"nilknow-jvm/rt"
)

type WIDE struct {
	modifiedInstruction base.Instruction
}

func (w *WIDE) FetchOperands(reader base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15:
		inst := &load.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		w.modifiedInstruction = inst
	case 0x16:
		inst := &load.LLOAD{}
		inst.Index = uint(reader.ReadUint16())
		w.modifiedInstruction = inst
	case 0x17:
		inst := &load.FLOAD{}
		inst.Index = uint(reader.ReadUint16())
		w.modifiedInstruction = inst
	case 0x18:
		inst := &load.DLOAD{}
		inst.Index = uint(reader.ReadUint16())
		w.modifiedInstruction = inst
	case 0x19:
		//todo aload
	case 0x36:
		inst := &store.ISTORE{}
		inst.Index = uint(reader.ReadUint16())
		w.modifiedInstruction = inst
	case 0x37:
		inst := &store.LSTORE{}
		inst.Index = uint(reader.ReadUint16())
		w.modifiedInstruction = inst
	case 0x38:
		// fstore
	case 0x39:
		// dstore
	case 0x3a:
		// astore
	case 0x84:
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadInt16())
		w.modifiedInstruction = inst

	case 0xa9: // ret
		panic("Unsupported opcode: 0xa9!")
	}
}

func (w *WIDE) Execute(frame rt.Frame) {
	w.modifiedInstruction.Execute(frame)
}

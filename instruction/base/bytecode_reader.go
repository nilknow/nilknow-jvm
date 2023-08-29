package base

type BytecodeReader struct {
	code []byte
	pc   int
}

func (r *BytecodeReader) PC() int {
	return r.pc
}

func (r *BytecodeReader) SetPc(pc int) {
	r.pc = pc
}

func (r *BytecodeReader) Reset(code []byte, pc int) {
	r.code = code
	r.pc = pc
}

func (r *BytecodeReader) ReadUint8() uint8 {
	i := r.code[r.pc]
	r.pc++
	return i
}

func (r *BytecodeReader) ReadInt8() int8 {
	i := r.code[r.pc]
	r.pc++
	return int8(i)
}

func (r *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(r.ReadUint8())
	byte2 := uint16(r.ReadUint8())
	return (byte1 << 8) | byte2
}

func (r *BytecodeReader) ReadInt16() int16 {
	byte1 := int16(r.ReadUint8())
	byte2 := int16(r.ReadUint8())
	return (byte1 << 8) | byte2
}

func (r *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(r.ReadUint8())
	byte2 := int32(r.ReadUint8())
	byte3 := int32(r.ReadUint8())
	byte4 := int32(r.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}

func (r *BytecodeReader) ReadInt32s(n int32) []int32 {
	ints := make([]int32, n)
	for i := range ints {
		ints[i] = r.ReadInt32()
	}
	return ints
}

func (r *BytecodeReader) SkipPadding() {
	for r.pc%4 != 0 {
		r.ReadUint8()
	}
}

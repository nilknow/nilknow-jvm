package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

func (r ClassReader) readConstantPool() ConstantPool {
	cpCount := int(r.readUint16())
	cp := make([]ConstantInfo, cpCount)

	for i := 1; i < cpCount; i++ {
		cp[i] = r.readConstantInfo(cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}


func (r ClassReader) readConstantInfo( cp ConstantPool) ConstantInfo {
	tag := r.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(r)
	return c
}


func (r ClassReader) readUint8() uint8 {
	val := r.data[0]
	r.data = r.data[1:]
	return val
}

func (r ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(r.data[0:2])
	r.data = r.data[2:]
	return val
}

func (r ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(r.data[0:4])
	r.data = r.data[4:]
	return val
}

func (r ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(r.data[0:8])
	r.data = r.data[8:]
	return val
}

func (r ClassReader) readBytes(length uint32) []byte {
	bytes := r.data[:length]
	r.data = r.data[length:]
	return bytes
}

func (r ClassReader) readUint16s() []uint16 {
	n := r.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = r.readUint16()
	}
	return s
}

package classfile

import (
	"encoding/binary"
	"fmt"
)

type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cf = &ClassFile{}

	cf.magic = readUnit32(&classData)
	if cf.magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic number is not correct!")
	}
	cf.minorVersion = readUnit16(&classData)
	cf.majorVersion = readUnit16(&classData)
	switch cf.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52, 53, 54, 55:
		if cf.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")

	return
}

func readUnit32(classData *[]byte) uint32 {
	val := binary.BigEndian.Uint32(*classData)
	*classData = (*classData)[4:]
	return val
}

func readUnit16(classData *[]byte) uint16 {
	val := binary.BigEndian.Uint16(*classData)
	*classData = (*classData)[2:]
	return val
}

func (cf *ClassFile) Magic() uint32 {
	return cf.magic
}

func (cf *ClassFile) MinorVersion() uint16 {
	return cf.minorVersion
}

func (cf *ClassFile) MajorVersion() uint16 {
	return cf.majorVersion
}

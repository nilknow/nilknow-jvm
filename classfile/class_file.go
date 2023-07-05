package classfile

import (
	"encoding/binary"
	"fmt"
)

type ClassFile struct {
	magic uint32
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

	magic := readUnit32(classData)
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic number is not correct!")
	}
	cf.magic =magic

	return
}

func readUnit32(classData []byte) uint32 {
	return binary.BigEndian.Uint32(classData)
}

func (cf *ClassFile) Magic() uint32 {
	return cf.magic
}


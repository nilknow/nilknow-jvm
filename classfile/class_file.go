package classfile

import (
	"fmt"
)

type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func (cf *ClassFile) Methods() []*MemberInfo {
	return cf.methods
}

func (cf *ClassFile) SetMethods(methods []*MemberInfo) {
	cf.methods = methods
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

	cr := &ClassReader{classData}
	cf = &ClassFile{}

	cf.magic = cr.readUint32()
	if cf.magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic number is not correct!")
	}
	cf.minorVersion = cr.readUint16()
	cf.majorVersion = cr.readUint16()
	switch cf.majorVersion {
	case 45:
	case 46, 47, 48, 49, 50, 51, 52, 53, 54, 55:
		if cf.minorVersion == 0 {
		}
	default:
		panic("java.lang.UnsupportedClassVersionError!")
	}
	cf.constantPool = cr.readConstantPool()
	cf.accessFlags = cr.readUint16()
	cf.thisClass = cr.readUint16()
	cf.superClass = cr.readUint16()
	cf.interfaces = cr.readUint16s()
	cf.methods = readMembers(cr, cf.constantPool)
	cf.attributes = readAttributes(cr, cf.constantPool)
	return
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

func (cf *ClassFile) ClassName() string {
	return cf.constantPool.getClassName(cf.thisClass)
}

func (cf *ClassFile) SuperClassName() string {
	if cf.superClass > 0 {
		return cf.constantPool.getClassName(cf.superClass)
	}
	return ""
}

func (cf *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(cf.interfaces))
	for i, cpIndex := range cf.interfaces {
		interfaceNames[i] = cf.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

func (cf *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

func (cf *ClassFile) readAndCheckVersion(reader *ClassReader) {
	cf.minorVersion = reader.readUint16()
	cf.majorVersion = reader.readUint16()
	switch cf.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if cf.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

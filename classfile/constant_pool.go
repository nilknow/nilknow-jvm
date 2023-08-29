package classfile

type ConstantPool []ConstantInfo

func (cp ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := cp[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

func (cp ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo:=cp.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name:=cp.getUtf8(ntInfo.nameIndex)
	_type := cp.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

func (cp ConstantPool) getClassName(index uint16) string {
	classInfo := cp.getConstantInfo(index).(*ConstantClassInfo)
	return cp.getUtf8(classInfo.nameIndex)
}

func (cp ConstantPool) getUtf8(index uint16) string {
	utf8Info := cp.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}

func (cp ConstantPool) getUint16(flags uint16) uint16 {
//	todo
	return uint16(1)
}

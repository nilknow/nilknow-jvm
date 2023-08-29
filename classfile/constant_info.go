package classfile

const (
	CONSTANT_CLASS              = 7
	CONSTANT_FIELDREF           = 9
	CONSTANT_METHODREF          = 10
	CONSTANT_INTERFACTMETHODREF = 11
	CONSTANT_STRING             = 8
	CONSTANT_INTEGER            = 3
	CONSTANT_FLOAT              = 4
	CONSTANT_LONG               = 5
	CONSTANT_DOUBLE             = 6
	CONSTANT_NAMEANDTYPE        = 12
	CONSTANT_UTF8               = 1
	CONSTANT_METHODHANDLE       = 15
	CONSTANT_METHODTYPE         = 16
	CONSTANT_INVOKEDYNAMIC      = 18
)

type ConstantInfo interface {
	readInfo(reader ClassReader)
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_INTEGER: return &ConstantIntegerInfo{}
	case CONSTANT_FLOAT: return &ConstantFloatInfo{}
	case CONSTANT_LONG: return &ConstantLongInfo{}
	case CONSTANT_DOUBLE: return &ConstantDoubleInfo{}
	case CONSTANT_UTF8: return &ConstantUtf8Info{}
	case CONSTANT_STRING: return &ConstantStringInfo{cp: cp}
	case CONSTANT_CLASS: return &ConstantClassInfo{cp: cp}
	case CONSTANT_FIELDREF:
		return &ConstantFieldRefInfo{ConstantMemberRefInfo{cp: cp}}
	case CONSTANT_METHODREF:
		return &ConstantMethodRefInfo{ConstantMemberRefInfo{cp: cp}}
	case CONSTANT_INTERFACTMETHODREF:
		return &ConstantInterfaceMethodRefInfo{ConstantMemberRefInfo{cp: cp}}
	case CONSTANT_NAMEANDTYPE: return &ConstantNameAndTypeInfo{}
	case CONSTANT_METHODTYPE: return &ConstantMethodTypeInfo{}
	case CONSTANT_METHODHANDLE: return &ConstantMethodHandleInfo{}
	case CONSTANT_INVOKEDYNAMIC: return &ConstantInvokeDynamicInfo{}
	default: panic("java.lang.ClassFormatError: constant pool tag!")
	}
}

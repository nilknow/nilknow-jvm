package classfile

type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}

func (mi *MemberInfo) AccessFlags() uint16 {
	return mi.cp.getUint16(mi.accessFlags)
}

func (mi *MemberInfo) Name() string {
	return mi.cp.getUtf8(mi.nameIndex)
}

func (mi *MemberInfo) Descriptor() string {
	return mi.cp.getUtf8(mi.descriptorIndex)
}

func (mi *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range mi.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}
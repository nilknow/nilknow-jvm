package classfile

type LocalVariableTableAttribute struct {
	localVariableTable []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc       uint16
	localVariable uint16
}

func (l *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	localVariableTableLength := reader.readUint16()
	l.localVariableTable = make([]*LocalVariableTableEntry, localVariableTableLength)
	for i := range l.localVariableTable {
		l.localVariableTable[i] = &LocalVariableTableEntry{
			startPc:       reader.readUint16(),
			localVariable: reader.readUint16(),
		}
	}
}

package classfile

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc uint16
	lineNumber uint16
}

func (l *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength:=reader.readUint16()
	l.lineNumberTable=make([]*LineNumberTableEntry,lineNumberTableLength)
	for i :=range l.lineNumberTable {
		l.lineNumberTable[i]=&LineNumberTableEntry{
			startPc: reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
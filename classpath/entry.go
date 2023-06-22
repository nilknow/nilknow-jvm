package classpath

import "os"
import "strings"

const pathListSeparator=string(os.PathListSeparator)

type Entry interface {
	readClass(className string)([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path,pathListSeparator){
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path,"*"){
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(strings.ToUpper(path),".JAR")||strings.HasSuffix(strings.ToUpper(path), ".ZIP") {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}



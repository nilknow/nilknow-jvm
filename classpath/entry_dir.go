package classpath

import (
	"os"
)
import "path/filepath"

type DirEntry struct{
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err :=filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &DirEntry{absDir}
}

func (entry *DirEntry) readClass(className string) ([]byte, Entry, error){
	fileName := filepath.Join(entry.absDir, className)
	data,err:=os.ReadFile(fileName)
	return data, entry, err
}

func (entry *DirEntry) String() string {
	return entry.absDir
}

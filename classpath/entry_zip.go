package classpath

import (
	"archive/zip"
	"errors"
	"io"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

//todo unzip every time readClass is called, need to optimize
func (entry *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	reader, err := zip.OpenReader(entry.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer reader.Close()

	for _, f := range reader.File {
		if f.Name == className {
			classReader, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer classReader.Close()

			data, err := io.ReadAll(classReader)
			if err != nil {
				return nil, nil, err
			}

			return data, entry, nil
		}
	}

	return nil, nil, errors.New("class not found: "+ className)
}

func (entry *ZipEntry) String() string {
	return entry.absPath
}

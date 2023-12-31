package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	var compositeEntry []Entry
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (entries CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range entries {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (entries CompositeEntry) String() string {
	strs := make([]string, len(entries))
	for i, entry := range entries {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}

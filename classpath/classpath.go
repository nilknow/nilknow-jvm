package classpath

import (
	"nilknow-jvm/file"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption string, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootClasspath(jreOption)
	cp.parseExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (cp *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := cp.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := cp.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return cp.userClasspath.readClass(className)
}

func (cp *Classpath) parseBootClasspath(jreOption string) {
	jreDir := file.GetJreDir(jreOption)
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	cp.bootClasspath = newWildcardEntry(jreLibPath)
}

func (cp *Classpath) parseExtClasspath(jreOption string) {
	jreDir := file.GetJreDir(jreOption)
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	cp.extClasspath = newWildcardEntry(jreExtPath)
}

func (cp *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	cp.userClasspath = newEntry(cpOption)
}

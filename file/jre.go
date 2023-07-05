package file

import (
	"os"
	"path/filepath"
)

func GetJreDir(jreOption string) string {
	if jreOption != "" {
		if exists(jreOption) {
			return jreOption
		}
		panic(jreOption + " doesn't exist")
	}

	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder")
}

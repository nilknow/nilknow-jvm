package file

import "os"

// todo what if the path is being used by another process?
func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

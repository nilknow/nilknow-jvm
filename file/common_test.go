package file

import "testing"

func TestExist(t *testing.T) {
	exist := exists(".")
	if !exist {
		t.Errorf("exists() should return ture for . as input")
	}

	exist = exists("file_not_exists")
	if exist {
		t.Errorf("exists() should return false for file_not_exists as input")
	}
}

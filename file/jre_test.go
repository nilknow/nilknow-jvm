package file

import "testing"

func TestGetJreDir(t *testing.T) {

	got := GetJreDir("path_not_exist")
	want := "not_empty_jre_path"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got = GetJreDir("./jre")
	want = "./jre"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

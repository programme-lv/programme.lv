package utils

import "testing"

func TestMakeTempDir(t *testing.T) {
	path, err := MakeTempDir()
	if err != nil {
		t.Error(err)
	}
	t.Log(path)
}

package utils

import (
	"os"
	"path/filepath"
)

func MakeTempDir() (path string, err error) {

	tempDirDirPath := filepath.Join("/tmp", "deikstra")
	err = os.MkdirAll(tempDirDirPath, os.ModePerm)
	if err != nil {
		return
	}

	return os.MkdirTemp(tempDirDirPath, "")
}

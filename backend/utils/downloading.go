package utils

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func DownloadMultiPartFile(file multipart.File, filename string) (downPath string, err error) {
	// CREATE TEMPORARY DIRECTORY
	tmpDir := filepath.Join("/tmp", "programme")
	_ = os.MkdirAll(tmpDir, os.ModePerm)
	tmpDir, _ = os.MkdirTemp(tmpDir, "")
	downPath = filepath.Join(tmpDir, filename)

	// DOWNLOAD FILE
	err = SaveMultiPartFile(file, downPath)
	return
}

func SaveMultiPartFile(file multipart.File, dst string) error {
	out, err := os.Create(dst)
	if err != nil {
		return err
	}

	_, err = io.Copy(out, file)
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}
	err = out.Close()
	if err != nil {
		return err
	}

	return nil
}

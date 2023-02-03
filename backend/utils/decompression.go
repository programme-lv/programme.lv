package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// DecompressZIP src is the path of .zip to decompress into dst
func DecompressZIP(zipPath string) (unzippedPath string, err error) {
	// make sure zipPath has .zip extension
	srcExtension := filepath.Ext(zipPath)
	if srcExtension != ".zip" {
		err = fmt.Errorf("file %s should have .zip extension", zipPath)
		return
	}

	// remove .zip extension from zipPath
	unzippedPath = zipPath[:len(zipPath)-len(".zip")]

	var r *zip.ReadCloser
	r, err = zip.OpenReader(zipPath)
	if err != nil {
		return
	}
	for _, f := range r.File {
		filePath := filepath.Join(unzippedPath, f.Name)

		if f.FileInfo().IsDir() {
			_ = os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return
		}

		var dstFile *os.File
		dstFile, err = os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return
		}

		var fileInArchive io.ReadCloser
		fileInArchive, err = f.Open()
		if err != nil {
			return
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			panic(err)
		}

		_ = dstFile.Close()
		_ = fileInArchive.Close()
	}
	_ = r.Close()
	return
}

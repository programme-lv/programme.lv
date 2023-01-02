package utils

import (
	"archive/zip"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

// DecompressZIP src is the path of .zip to decompress into dst
func DecompressZIP(src string, dst string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range r.File {
		filePath := filepath.Join(dst, f.Name)
		log.Println("unzipping file ", filePath)

		if f.FileInfo().IsDir() {
			_ = os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return err
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		fileInArchive, err := f.Open()
		if err != nil {
			panic(err)
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			panic(err)
		}

		_ = dstFile.Close()
		_ = fileInArchive.Close()
	}
	_ = r.Close()
	return nil
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

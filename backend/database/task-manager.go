package database

import (
	"archive/zip"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

type TaskManager struct {
}

func CreateTaskManager() *TaskManager {
	return &TaskManager{}
}

// CreateTask creates the task, validates it, names it
func (tm *TaskManager) CreateTask(taskFile multipart.File) error {
	dirPath := filepath.Join("/tmp", "deikstra")
	_ = os.MkdirAll(dirPath, os.ModePerm)
	tmpDir, _ := os.MkdirTemp(dirPath, "")
	download := filepath.Join(tmpDir, "download")

	out, err := os.Create(download)
	if err != nil {
		return err
	}

	_, err = io.Copy(out, taskFile)
	if err != nil {
		return err
	}

	err = taskFile.Close()
	if err != nil {
		return err
	}
	err = out.Close()
	if err != nil {
		return err
	}

	r, err := zip.OpenReader(download)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range r.File {
		filePath := filepath.Join(tmpDir, "unzipped", f.Name)
		log.Println("unzipping file ", filePath)

		if f.FileInfo().IsDir() {
			os.MkdirAll(filePath, os.ModePerm)
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

		dstFile.Close()
		fileInArchive.Close()
	}
	_ = r.Close()
	return nil
}

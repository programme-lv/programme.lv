package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type Executable struct {
	srcPath string
	exePath string
}

func (e *Executable) Execute(stdin io.ReadCloser) (stdout string, stderr string, err error) {
	cmd := exec.Command(e.exePath)
	stdoutPipe, _ := cmd.StdoutPipe()
	stderrPipe, _ := cmd.StderrPipe()
	if err = cmd.Start(); err != nil {
		return
	}
	stdoutBytes, _ := ioutil.ReadAll(stdoutPipe)
	stderrBytes, _ := ioutil.ReadAll(stderrPipe)
	stdout = string(stdoutBytes)
	stderr = string(stderrBytes)
	err = cmd.Wait()
	if err != nil {
		return
	}
	return
}

func CreateExecutable(srcCode string, langId string) (*Executable, error) {
	exe := &Executable{}

	exeDir, err := os.MkdirTemp("/tmp/deikstra/", "")
	if err != nil {
		return exe, err
	}

	srcFile, nil := os.Create(filepath.Join(exeDir, "main.cpp"))
	if err != nil {
		return exe, err
	}
	defer srcFile.Close()

	exe.srcPath = srcFile.Name()

	_, err = srcFile.WriteString(srcCode)
	if err != nil {
		return exe, err
	}

	exeFile := filepath.Join(exeDir, "exe")
	cmd := exec.Command("g++", exe.srcPath, "-o", exeFile)

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		return exe, nil
	}

	stdoutStr, _ := ioutil.ReadAll(stdout)
	stderrStr, _ := ioutil.ReadAll(stderr)
	if err := cmd.Wait(); err != nil {
		log.Printf("stdout: %v\n", string(stdoutStr))
		log.Printf("stderr: %v\n", string(stderrStr))
		log.Printf("cmd wait err: %v\n", err)
		return exe, nil
	}
	log.Printf("stdout: %v\n", string(stdoutStr))
	log.Printf("stderr: %v\n", string(stderrStr))
	// TODO: figure out what to do if err is nil but stderr isn't empty

	exe.exePath = exeFile

	return exe, nil
}

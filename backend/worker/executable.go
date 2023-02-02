package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type Executable struct {
	srcPath string
	exePath string
}

type ExecutionResult struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

func (e *Executable) Execute(stdin io.Reader) (exeRes ExecutionResult, err error) {
	cmd := exec.Command(e.exePath)
	stdinPipe, _ := cmd.StdinPipe()
	stdoutPipe, _ := cmd.StdoutPipe()
	stderrPipe, _ := cmd.StderrPipe()
	_, err = io.Copy(stdinPipe, stdin)
	if err != nil {
		return
	}
	if err = cmd.Start(); err != nil {
		return
	}
	stdoutBytes, _ := io.ReadAll(stdoutPipe)
	stderrBytes, _ := io.ReadAll(stderrPipe)
	err = cmd.Wait()
	if err != nil {
		return
	}
	exeRes.Stdout = string(stdoutBytes)
	exeRes.Stderr = string(stderrBytes)
	exeRes.ExitCode = cmd.ProcessState.ExitCode()
	return
}

func NewExecutable(srcCode string, langId string) (exe *Executable, compilationExeRes ExecutionResult, err error) {
	exe = &Executable{}

	dirPath := filepath.Join("/tmp", "deikstra")
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return
	}

	exeDir, err := os.MkdirTemp(dirPath, "")
	if err != nil {
		return
	}

	srcFile, nil := os.Create(filepath.Join(exeDir, "main.cpp"))
	if err != nil {
		return
	}
	defer func(srcFile *os.File) {
		err := srcFile.Close()
		if err != nil {

		}
	}(srcFile)

	exe.srcPath = srcFile.Name()

	_, err = srcFile.WriteString(srcCode)
	if err != nil {
		return
	}

	exeFile := filepath.Join(exeDir, "exe")
	cmd := exec.Command("g++", exe.srcPath, "-o", exeFile)

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	if err = cmd.Start(); err != nil {
		return
	}

	stdoutStr, _ := io.ReadAll(stdout)
	stderrStr, _ := io.ReadAll(stderr)
	if err = cmd.Wait(); err != nil {
		log.Printf("stdout: %v\n", string(stdoutStr))
		log.Printf("stderr: %v\n", string(stderrStr))
		log.Printf("cmd wait err: %v\n", err)
		return
	}
	compilationExeRes.Stdout = string(stdoutStr)
	compilationExeRes.Stderr = string(stderrStr)
	compilationExeRes.ExitCode = cmd.ProcessState.ExitCode()

	exe.exePath = exeFile

	return
}

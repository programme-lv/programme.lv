package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type Executable struct {
	exePath string
}

type ExecutionResult struct {
	Stdout   string
	Stderr   string
	ExitCode int
	CpuTime  int
	WallTime int
	MemUsage int
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

	var exeDir string
	exeDir, err = makeTmpDir()
	if err != nil {
		return
	}

	switch langId {
	case "C++17":
		exe.exePath, _, err = compileCpp(srcCode, exeDir)
		return
	}
	// create an error if langId is not supported
	err = fmt.Errorf("langId \"%s\" is not supported", langId)
	return
}

func makeTmpDir() (path string, err error) {

	dirPath := filepath.Join("/tmp", "deikstra")
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return
	}

	return os.MkdirTemp(dirPath, "")
}

func compileCpp(code string, dir string) (exePath string, execRes ExecutionResult, err error) {

	srcFile, nil := os.Create(filepath.Join(dir, "main.cpp"))
	if err != nil {
		return
	}
	defer func(srcFile *os.File) {
		err := srcFile.Close()
		if err != nil {

		}
	}(srcFile)

	_, err = srcFile.WriteString(code)
	if err != nil {
		return
	}

	exePath = filepath.Join(dir, "exe")
	cmd := exec.Command("g++", srcFile.Name(), "-o", exePath)

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
	execRes.Stdout = string(stdoutStr)
	execRes.Stderr = string(stderrStr)
	execRes.ExitCode = cmd.ProcessState.ExitCode()

	return
}

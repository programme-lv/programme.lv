package execution

import (
	"fmt"
	"github.com/KrisjanisP/deikstra/service/worker/utils"
	"io"
	"os/exec"
	"path/filepath"
	"strings"
)

type Result struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

func Execute(cmd *exec.Cmd, stdin io.Reader) (res Result, err error) {
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
	res.Stdout = string(stdoutBytes)
	res.Stderr = string(stderrBytes)
	res.ExitCode = cmd.ProcessState.ExitCode()
	return
}

func ExecuteCmd(fullCmd string, stdin io.Reader) (executionRes Result, err error) {
	cmdName := strings.Fields(fullCmd)[0]
	cmdArgs := strings.Fields(fullCmd)[1:]
	cmd := exec.Command(cmdName, cmdArgs...)
	return Execute(cmd, stdin)
}

type ExpandedResult struct {
	Result
	CpuTime  int
	WallTime int
	MemUsage int
}

type Constraints struct {
	memoryLimit   *int // kilobytes
	timeLimit     *int // seconds
	wallTimeLimit *int // seconds
	extraTime     *int // seconds
}

func (box *IsolateBox) Execute(boxCmd string, stdin io.Reader, constraints Constraints) (res ExpandedResult, err error) {
	var directory string
	directory, err = utils.MakeTempDir()
	if err != nil {
		return
	}

	metaFilePath := filepath.Join(directory, "meta.txt")

	cmdArgs := make([]string, 0)

	if constraints.memoryLimit != nil {
		cmdArgs = append(cmdArgs, fmt.Sprintf("--mem=%d", *constraints.memoryLimit))
	}
	if constraints.timeLimit != nil {
		cmdArgs = append(cmdArgs, fmt.Sprintf("--time=%d", *constraints.timeLimit))
	}
	if constraints.wallTimeLimit != nil {
		cmdArgs = append(cmdArgs, fmt.Sprintf("--wall-time=%d", *constraints.wallTimeLimit))
	}
	if constraints.extraTime != nil {
		cmdArgs = append(cmdArgs, fmt.Sprintf("--extra-time=%d", *constraints.extraTime))
	}

	cmdArgs = append(cmdArgs, fmt.Sprintf("--meta=%s", metaFilePath))
	cmdArgs = append(cmdArgs, fmt.Sprintf("--box-id=%d", box.id))
	cmdArgs = append(cmdArgs, "--run", boxCmd)
	cmd := exec.Command("isolate", cmdArgs...)

	var executionRes Result
	executionRes, err = Execute(cmd, stdin)
	if err != nil {
		return
	}

	res.Stdout = executionRes.Stdout
	res.Stderr = executionRes.Stderr
	res.ExitCode = executionRes.ExitCode

	// TODO: read meta file
	return
}

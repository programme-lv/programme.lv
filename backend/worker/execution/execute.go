package execution

import (
	"fmt"
	"github.com/KrisjanisP/deikstra/service/worker/utils"
	"io"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

type Result struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

func Execute(cmd *exec.Cmd, stdin io.ReadCloser) (res Result, err error) {
	stdinPipe, _ := cmd.StdinPipe()
	stdoutPipe, _ := cmd.StdoutPipe()
	stderrPipe, _ := cmd.StderrPipe()
	if stdin != nil {
		_, err = io.Copy(stdinPipe, stdin)
		_ = stdin.Close()
	}
	_ = stdinPipe.Close()
	if err != nil {
		return
	}
	if err = cmd.Start(); err != nil {
		return
	}
	stdoutBytes, _ := io.ReadAll(stdoutPipe)
	stderrBytes, _ := io.ReadAll(stderrPipe)
	err = cmd.Wait()
	_ = stdoutPipe.Close()
	_ = stderrPipe.Close()
	res.Stdout = string(stdoutBytes)
	res.Stderr = string(stderrBytes)
	res.ExitCode = cmd.ProcessState.ExitCode()
	return
}

func ExecuteCmd(fullCmd string, stdin io.ReadCloser) (executionRes Result, err error) {
	cmdName := strings.Fields(fullCmd)[0]
	cmdArgs := strings.Fields(fullCmd)[1:]
	cmd := exec.Command(cmdName, cmdArgs...)
	return Execute(cmd, stdin)
}

type ExtendedResult struct {
	Result
	CpuTime  int
	WallTime int
	MemUsage int
}

type Constraints struct {
	memoryLimit   int // kilobytes
	timeLimit     int // seconds
	wallTimeLimit int // seconds
	extraTime     int // seconds
}

var DefaultConstraints = Constraints{
	memoryLimit:   256 * 1024,
	timeLimit:     3,
	wallTimeLimit: 10,
	extraTime:     2,
}

func (box *IsolateBox) Execute(boxCmd string, stdin io.ReadCloser, constraints Constraints) (res *ExtendedResult, err error) {
	res = &ExtendedResult{}

	var directory string
	directory, err = utils.MakeTempDir()
	if err != nil {
		return
	}

	metaFilePath := filepath.Join(directory, "meta.txt")

	cmdArgs := make([]string, 0)

	cmdArgs = append(cmdArgs, fmt.Sprintf("--mem=%d", constraints.memoryLimit))
	cmdArgs = append(cmdArgs, fmt.Sprintf("--time=%d", constraints.timeLimit))
	cmdArgs = append(cmdArgs, fmt.Sprintf("--wall-time=%d", constraints.wallTimeLimit))
	cmdArgs = append(cmdArgs, fmt.Sprintf("--extra-time=%d", constraints.extraTime))

	cmdArgs = append(cmdArgs, fmt.Sprintf("--meta=%s", metaFilePath))
	cmdArgs = append(cmdArgs, fmt.Sprintf("--box-id=%d", box.id))
	cmdArgs = append(cmdArgs, "--cg")
	cmdArgs = append(cmdArgs, "--processes=4")
	cmdArgs = append(cmdArgs, "--env=PATH=/usr/bin")
	cmdArgs = append(cmdArgs, "--run", "--")
	for _, boxArg := range strings.Split(boxCmd, " ") {
		cmdArgs = append(cmdArgs, boxArg)
	}
	cmd := exec.Command("isolate", cmdArgs...)
	cmd.Dir = box.BoxPath

	log.Println(strings.Join(strings.Split(cmd.String(), " "), ","))

	var executionRes Result
	executionRes, err = Execute(cmd, stdin)
	res.Stdout = executionRes.Stdout
	res.Stderr = executionRes.Stderr
	res.ExitCode = executionRes.ExitCode
	if err != nil {
		return
	}

	// TODO: read meta file
	return
}

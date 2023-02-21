package execution

import (
	"io"
	"strings"
	"testing"
)

func TestExecuteCmd(t *testing.T) {
	res, err := ExecuteCmd("echo test", nil)
	if err != nil {
		t.Error(err)
	}
	if strings.Trim(res.Stdout, " \n") != "test" {
		t.Error("Wrong stdout, received: " + "\"" + res.Stdout + "\"")
	}
	if res.Stderr != "" {
		t.Error("Wrong stderr")
	}
	if res.ExitCode != 0 {
		t.Error("Wrong exit code")
	}
}

func TestExecuteCmdWithStdin(t *testing.T) {
	res, err := ExecuteCmd("cat", io.NopCloser(strings.NewReader("test")))
	if err != nil {
		t.Error(err)
	}
	if strings.Trim(res.Stdout, " \n") != "test" {
		t.Error("Wrong stdout, received: " + "\"" + strings.Trim(res.Stdout, " \n") + "\"")
	}
	if res.Stderr != "" {
		t.Error("Wrong stderr")
	}
	if res.ExitCode != 0 {
		t.Error("Wrong exit code")
	}
}

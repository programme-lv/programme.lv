package executable

import (
	"github.com/KrisjanisP/deikstra/service/models"
	"github.com/KrisjanisP/deikstra/service/worker/execution"
	"testing"
)

func TestNewIsolatedExecutable(t *testing.T) {
	ic := execution.NewIsolateController(4)
	box, err := ic.NewIsolateBox()
	if err != nil {
		t.Error(err)
	}
	defer box.Release()

	t.Log("box path: ", box.BoxPath)
	srcCode := &SrcCode{
		code:     "print('Hello World!')",
		language: models.Language{Filename: "main.py", CompileCmd: "python3 -m main.py", ExecuteCmd: "python3 main.py"},
	}

	ie, compRes, err := NewIsolatedExecutable(box, srcCode)
	if err != nil {
		t.Error(err)
	}
	if ie == nil {
		t.Error("IsolatedExecutable is nil")
	}
	t.Log(compRes)
}

package executable

import (
	"github.com/KrisjanisP/deikstra/service/models"
	"github.com/KrisjanisP/deikstra/service/worker/execution"
	"io"
	"os"
	"path/filepath"
)

type IsolatedExecutable struct {
	box *execution.IsolateBox
}

type SrcCode struct {
	code     string
	language models.Language
}

func NewIsolatedExecutable(box *execution.IsolateBox, srcCode *SrcCode) (exe *IsolatedExecutable, compilation *execution.ExtendedResult, err error) {
	// place srcCode in the box
	_, err = os.Create(filepath.Join(box.BoxPath, srcCode.language.Filename))
	if err != nil {
		return
	}
	// compile the executable in the box

	exe = &IsolatedExecutable{box: box}
	return
}

func (e *IsolatedExecutable) Execute(stdin io.Reader) (res execution.ExtendedResult, err error) {
	// execute the executable in the box
	// read stdout and stderr from the box
	// return the result
	return
}

package executable

import (
	"github.com/KrisjanisP/deikstra/service/models"
	"github.com/KrisjanisP/deikstra/service/worker/execution"
	"io"
	"os"
	"path/filepath"
)

type IsolatedExecutable struct {
	box     *execution.IsolateBox
	srcCode *SrcCode
}

type SrcCode struct {
	code     string
	language models.Language
}

func NewIsolatedExecutable(box *execution.IsolateBox, srcCode *SrcCode) (exe *IsolatedExecutable, compilation *execution.ExtendedResult, err error) {
	// place source code in the box
	srcFilePath := filepath.Join(box.BoxPath, srcCode.language.Filename)
	var srcFile *os.File
	srcFile, err = os.Create(srcFilePath)
	if err != nil {
		return
	}
	_, err = srcFile.Write([]byte(srcCode.code))
	if err != nil {
		return
	}

	// compile the executable in the box
	if srcCode.language.CompileCmd != nil {
		compilation, err = box.Execute(*srcCode.language.CompileCmd, nil, execution.DefaultConstraints)
		if err != nil {
			return
		}
	}

	exe = &IsolatedExecutable{box: box, srcCode: srcCode}
	return
}

func (e *IsolatedExecutable) Execute(stdin io.ReadCloser) (res *execution.ExtendedResult, err error) {
	res, err = e.box.Execute(e.srcCode.language.ExecuteCmd, stdin, execution.DefaultConstraints)
	return
}

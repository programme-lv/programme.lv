package execution

import (
	"fmt"
	"path/filepath"
	"strings"
)

type IsolateController struct {
	boxIds chan int
}

type IsolateBox struct {
	id      int
	BoxPath string
	ic      *IsolateController
}

func (box *IsolateBox) Release() {
	box.ic.boxIds <- box.id
}

func (ic *IsolateController) NewIsolateBox() (box *IsolateBox, err error) {
	// run isolate --cleaup --box-id {id}
	// run isolate --init --box-id {id} <-- returns path to box
	box = &IsolateBox{ic: ic}
	box.id = <-ic.boxIds

	_, err = ExecuteCmd(fmt.Sprintf("isolate --cleanup --box-id %d", box.id), nil)
	if err != nil {
		return
	}

	var initRes Result
	initRes, err = ExecuteCmd(fmt.Sprintf("isolate --init --box-id %d", box.id), nil)
	if err != nil {
		return
	}

	box.BoxPath = filepath.Join(strings.Trim(initRes.Stdout, " \n\r"), "box")

	return
}

func NewIsolateController(boxes int) *IsolateController {
	result := &IsolateController{boxIds: make(chan int, boxes)}
	for i := 0; i < boxes; i++ {
		result.boxIds <- i
	}
	return result
}

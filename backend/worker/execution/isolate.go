package execution

import "fmt"

type IsolateController struct {
	boxIds chan int
}

type IsolateBox struct {
	id      int
	boxPath string
}

func (ic *IsolateController) NewIsolateBox() (box *IsolateBox, err error) {
	// run isolate --cleaup --box-id {id}
	// run isolate --init --box-id {id} <-- returns path to box

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

	box.boxPath = initRes.Stdout

	return
}

func (ic *IsolateController) ReleaseIsolateBox(box *IsolateBox) {
	ic.boxIds <- box.id
}

func NewIsolateController(boxes int) *IsolateController {
	result := &IsolateController{boxIds: make(chan int, boxes)}
	for i := 0; i < boxes; i++ {
		result.boxIds <- i
	}
	return result
}

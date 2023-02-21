package execution

import "testing"

func TestNewIsolateController(t *testing.T) {
	ic := NewIsolateController(2)
	if ic == nil {
		t.Error("IsolateController is nil")
	}
}

func TestIsolateController_NewIsolateBox(t *testing.T) {
	boxes := 2
	ic := NewIsolateController(boxes)
	for i := 0; i < boxes; i++ {
		box, err := ic.NewIsolateBox()
		if err != nil {
			t.Error(err)
		}
		t.Log(box.BoxPath)
	}
}

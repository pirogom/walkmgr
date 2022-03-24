package walkmgr

import (
	"testing"
)

func TestWin(t *testing.T) {
	wm := NewWin("기본윈도", 640, 480, LAYOUT_VERT)

	if wm == nil {
		return
	}
	wm.Start()
}

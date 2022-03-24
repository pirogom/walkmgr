package walkmgr

import (
	"testing"
)

func TestDefaultWin(t *testing.T) {
	wm := NewWin("기본윈도", 640, 480, LAYOUT_VERT)
	wm.Start()
}

func TestNoMinBox(t *testing.T) {
	wm := NewWin("기본윈도-최소화X", 640, 480, LAYOUT_VERT)
	wm.DisableMinBox()

	wm.Start()
}

func TestNoMaxBox(t *testing.T) {
	wm := NewWin("기본윈도-최대화X", 640, 480, LAYOUT_VERT)
	wm.DisableMaxBox()

	wm.Start()
}

func TestNoSysMenu(t *testing.T) {
	wm := NewWin("기본윈도-시스템메뉴X", 640, 480, LAYOUT_VERT)
	wm.DisableSysmenu()

	wm.Start()
}

func TestNoTitle(t *testing.T) {
	wm := NewWin("기본윈도-타이틀바X", 640, 480, LAYOUT_VERT)
	wm.DisableTitleBar()

	wm.Start()
}

func TestNoResize(t *testing.T) {
	wm := NewWin("기본윈도-사이즈변경X", 640, 480, LAYOUT_VERT)
	wm.NoResize()
	wm.Start()
}


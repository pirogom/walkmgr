package walkmgr

import (
	"container/list"

	"github.com/pirogom/walk"
	"github.com/pirogom/win"
)

/**
*	walkmgr
**/
type walkmgr struct {
	window     *walk.MainWindow
	parentList *list.List
}

/**
*	NewWin
**/
func NewWin(title string, width int, height int, lt layoutType) *walkmgr {
	wm := walkmgr{}

	// set config
	win, winErr := walk.NewMainWindowWithName(defaultWinName)

	if winErr != nil {
		return nil
	}
	wm.window = win
	wm.parentList = list.New()

	wm.window.SetTitle(title)
	wm.window.SetWidth(width)
	wm.window.SetHeight(height)

	// set layout
	switch lt {
	case LAYOUT_VERT:
		wm.window.SetLayout(walk.NewVBoxLayout())
	case LAYOUT_HORI:
		wm.window.SetLayout(walk.NewHBoxLayout())
	default:
		wm.window.SetLayout(walk.NewVBoxLayout())
	}

	if defaultIcon != nil {
		wm.window.SetIcon(defaultIcon)
	}

	return &wm
}

/**
*	GetHWND
**/
func (wm *walkmgr) GetHWND() win.HWND {
	return wm.window.Handle()
}

/**
*	GetWindow
**/
func (wm *walkmgr) GetWindow() *walk.MainWindow {
	return wm.window
}

/**
*	SetTitle
**/
func (wm *walkmgr) SetTitle(title string) *walkmgr {
	wm.window.SetTitle(title)
	return wm
}

/**
*	SetSize
**/
func (wm *walkmgr) SetSize(width int, height int) *walkmgr {
	wm.window.SetWidth(width)
	wm.window.SetHeight(height)
	return wm
}

/**
*	SetMinSize
**/
func (wm *walkmgr) SetMinSize(width int, height int) *walkmgr {
	maxSize := wm.window.MaxSize()
	wm.window.SetMinMaxSize(walk.Size{Width: width, Height: height}, maxSize)
	return wm
}

/**
*	SetMaxSize
**/
func (wm *walkmgr) SetMaxSize(width int, height int) *walkmgr {
	minSize := wm.window.MinSize()
	wm.window.SetMinMaxSize(minSize, walk.Size{Width: width, Height: height})
	return wm
}

/**
* NoResize
**/
func (wm *walkmgr) NoResize() *walkmgr {
	defStyle := win.GetWindowLong(wm.GetHWND(), win.GWL_STYLE)
	newStyle := defStyle &^ win.WS_THICKFRAME
	win.SetWindowLong(wm.GetHWND(), win.GWL_STYLE, newStyle)
	return wm
}

/**
*	DisableMinBox
**/
func (wm *walkmgr) DisableMinBox() *walkmgr {
	defStyle := win.GetWindowLong(wm.GetHWND(), win.GWL_STYLE)
	newStyle := defStyle &^ win.WS_MINIMIZEBOX
	win.SetWindowLong(wm.GetHWND(), win.GWL_STYLE, newStyle)
	return wm
}

/**
*	DisableMaxBox
**/
func (wm *walkmgr) DisableMaxBox() *walkmgr {
	defStyle := win.GetWindowLong(wm.GetHWND(), win.GWL_STYLE)
	newStyle := defStyle &^ win.WS_MAXIMIZEBOX
	win.SetWindowLong(wm.GetHWND(), win.GWL_STYLE, newStyle)
	return wm
}

/**
*	DisableMinMaxBox
**/
func (wm *walkmgr) DisableMinMaxBox() *walkmgr {
	defStyle := win.GetWindowLong(wm.GetHWND(), win.GWL_STYLE)
	newStyle := defStyle &^ win.WS_MINIMIZEBOX
	newStyle = newStyle &^ win.WS_MAXIMIZEBOX
	win.SetWindowLong(wm.GetHWND(), win.GWL_STYLE, newStyle)
	return wm
}

/**
*	DisableCloseBox
**/
func (wm *walkmgr) DisableCloseBox() *walkmgr {
	defStyle := win.GetWindowLong(wm.GetHWND(), win.GWL_STYLE)
	newStyle := defStyle &^ win.WS_SYSMENU
	win.SetWindowLong(wm.GetHWND(), win.GWL_STYLE, newStyle)
	return wm
}

/**
*	DisableTitleBar
**/
func (wm *walkmgr) DisableTitleBar() *walkmgr {
	defStyle := win.GetWindowLong(wm.GetHWND(), win.GWL_STYLE)
	newStyle := defStyle &^ win.WS_CAPTION
	win.SetWindowLong(wm.GetHWND(), win.GWL_STYLE, newStyle)
	return wm
}

/**
*	SetForeground
**/
func (wm *walkmgr) SetForeground() {
	win.SetForegroundWindow(wm.GetHWND())
}

/**
*	Start
**/
func (wm *walkmgr) Start() {
	wm.window.Show()
	wm.window.Run()
}

/**
*	StartForeground
**/
func (wm *walkmgr) StartForeground() {
	wm.SetForeground()
	wm.window.Show()
	wm.window.Run()
}

/**
*	HideStart
**/
func (wm *walkmgr) HideStart() {
	wm.window.Hide()
	wm.window.Run()
}

/**
*	Hide
**/
func (wm *walkmgr) Hide() {
	wm.window.Hide()
}

/**
*	Show
**/
func (wm *walkmgr) Show() {
	wm.window.Show()
}

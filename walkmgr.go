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
func NewWin(title string, width int, height int, lt ...LayoutType) *walkmgr {
	wm := walkmgr{}

	// set config
	cfg := walk.MainWindowCfg{}
	cfg.Name = defaultWinName
	cfg.Bounds.SetLocation(walk.Point{X: win.CW_USEDEFAULT, Y: win.CW_USEDEFAULT})
	cfg.Bounds.SetSize(walk.Size{Width: width, Height: height})

	win, winErr := walk.NewMainWindowWithCfg(&cfg)

	if winErr != nil {
		panic("create window failed. please check manifest and .syso")
	}
	wm.window = win
	wm.parentList = list.New()

	wm.window.SetTitle(title)
	wm.window.SetWidth(width)
	wm.window.SetHeight(height)

	if len(lt) > 0 {
		// set layout
		switch lt[0] {
		case LAYOUT_VERT:
			wm.window.SetLayout(walk.NewVBoxLayout())
		case LAYOUT_HORI:
			wm.window.SetLayout(walk.NewHBoxLayout())
		default:
			wm.window.SetLayout(walk.NewVBoxLayout())
		}
	} else {
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
*	Window
**/
func (wm *walkmgr) Window() *walk.MainWindow {
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
*	DisableSysmenu
**/
func (wm *walkmgr) DisableSysmenu() *walkmgr {
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

/**
*	IgnoreClosing
**/
func (wm *walkmgr) IgnoreClosing() {
	wm.window.Closing().Attach(func(canceled *bool, reason walk.CloseReason) {
		if wm.window.Visible() {
			*canceled = true
		}
	})
}

/**
*	Sync
**/
func (wm *walkmgr) Sync(syncFunc func()) {
	wm.window.Synchronize(syncFunc)
}

/**
*	ForceClose
**/
func (wm *walkmgr) ForceClose() {
	wm.Sync(func() {
		wm.window.SetVisible(false)
		wm.window.Close()
	})
}

/**
* Append
**/
func (wm *walkmgr) Append(item walk.Widget) {
	if wm.parentList.Len() == 0 {
		wm.window.Children().Add(item)
	} else {
		parent := wm.parentList.Back().Value.(walk.Container)
		parent.Children().Add(item)
	}
}

/**
*	Parent
**/
func (wm *walkmgr) Parent() walk.Container {
	if wm.parentList.Len() > 0 {
		parent := wm.parentList.Back().Value.(walk.Container)
		return parent
	} else {
		return wm.window
	}
}

/**
*	Split
**/
func (wm *walkmgr) Split(lt ...LayoutType) (*walkmgr, *walk.Splitter) {
	var hs *walk.Splitter

	if len(lt) == 0 {
		hs, _ = walk.NewHSplitter(wm.Parent())
	} else {
		switch lt[0] {
		case LAYOUT_VERT:
			hs, _ = walk.NewVSplitter(wm.Parent())
		case LAYOUT_HORI:
			hs, _ = walk.NewHSplitter(wm.Parent())
		default:
			hs, _ = walk.NewHSplitter(wm.Parent())
		}
	}
	wm.parentList.PushBack(hs)
	return wm, hs
}

/**
*	GroupBox
**/
func (wm *walkmgr) GroupBox(title string, lt ...LayoutType) (*walkmgr, *walk.GroupBox) {
	gb, _ := walk.NewGroupBox(wm.Parent())
	gb.SetTitle(title)

	if len(lt) == 0 {
		gb.SetLayout(walk.NewVBoxLayout())
	} else {
		switch lt[0] {
		case LAYOUT_VERT:
			gb.SetLayout(walk.NewVBoxLayout())
		case LAYOUT_HORI:
			gb.SetLayout(walk.NewHBoxLayout())
		default:
			gb.SetLayout(walk.NewVBoxLayout())
		}
	}
	wm.parentList.PushBack(gb)
	return wm, gb
}

/**
*	End
**/
func (wm *walkmgr) End() {
	if wm.parentList.Len() > 0 {
		popData := wm.parentList.Remove(wm.parentList.Back())
		parent := wm.Parent()
		parent.Children().Add(popData.(walk.Widget))
	}
}

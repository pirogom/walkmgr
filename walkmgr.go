package walkmgr

import (
	"container/list"

	"github.com/pirogom/walk"
	"github.com/pirogom/win"
)

type WinCloseFunc func() bool
type WinStartFunc func()

/**
*	WalkUI
**/
type WalkUI struct {
	window        *walk.MainWindow
	parentList    *list.List
	IsIgnoreClose bool
	startingFunc  WinStartFunc
	closingFunc   WinCloseFunc
}

/**
*	NewWin
**/
func NewWin(title string, width int, height int, lt ...LayoutType) *WalkUI {
	wm := WalkUI{}

	// set config
	cfg := walk.MainWindowCfg{}
	cfg.Name = defaultWinName
	cfg.Bounds.SetLocation(walk.Point{X: win.CW_USEDEFAULT, Y: win.CW_USEDEFAULT})
	cfg.Bounds.SetSize(walk.Size{Width: width, Height: height})

	window, winErr := walk.NewMainWindowWithCfg(&cfg)

	if winErr != nil {
		panic("create window failed. please check manifest and .syso")
	}
	wm.window = window
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
		case LAYOUT_FLOW:
			wm.window.SetLayout(walk.NewFlowLayout())
		default:
			wm.window.SetLayout(walk.NewVBoxLayout())
		}
	} else {
		wm.window.SetLayout(walk.NewVBoxLayout())
	}

	if defaultIcon != nil {
		wm.window.SetIcon(defaultIcon)
	}

	// windows start
	wm.window.Starting().Attach(func() {
		wm.center()
		if wm.startingFunc != nil {
			wm.startingFunc()
		}
	})

	// closing
	wm.window.Closing().Attach(func(canceled *bool, reason walk.CloseReason) {
		if wm.IsIgnoreClose && wm.Window().Visible() {
			*canceled = true
			return
		}

		if wm.closingFunc != nil {
			if !wm.closingFunc() {
				*canceled = true
				return
			}
		}
	})

	return &wm
}

/**
*	NewFixed
**/
func NewFixed(title string, width int, height int, lt ...LayoutType) *WalkUI {
	wm := NewWin(title, width, height, lt...)
	wm.NoResize().DisableMinMaxBox()
	return wm
}

/**
*	NewAds
**/
func NewAds(title string, width int, height int) *WalkUI {
	wm := WalkUI{}

	// set config
	cfg := walk.MainWindowCfg{}
	cfg.Name = defaultWinName
	cfg.Bounds.SetLocation(walk.Point{X: win.CW_USEDEFAULT, Y: win.CW_USEDEFAULT})
	cfg.Bounds.SetSize(walk.Size{Width: width, Height: height})

	window, winErr := walk.NewMainWindowWithCfg(&cfg)

	if winErr != nil {
		panic("create window failed. please check manifest and .syso")
	}
	wm.window = window
	wm.parentList = list.New()

	wm.window.SetTitle(title)
	layout := walk.NewVBoxLayout()
	margin := walk.Margins{0, 0, 0, 0}
	layout.SetMargins(margin)
	layout.SetSpacing(0)
	wm.window.SetLayout(layout)
	wm.window.SetWidth(width)
	wm.window.SetHeight(height)
	wm.window.SetMinMaxSize(walk.Size{Width: width, Height: height}, walk.Size{Width: width, Height: height})

	wm.NoResize()
	wm.DisableTitleBar()

	// starting
	wm.window.Starting().Attach(func() {
		wm.adsPosition()

		if wm.startingFunc != nil {
			wm.startingFunc()
		}
	})

	// closing
	wm.window.Closing().Attach(func(canceled *bool, reason walk.CloseReason) {
		if wm.IsIgnoreClose && wm.Window().Visible() {
			*canceled = true
			return
		}

		if wm.closingFunc != nil {
			if !wm.closingFunc() {
				*canceled = true
				return
			}
		}
	})

	return &wm
}

/**
*	adsPosition
**/
func (wm *WalkUI) adsPosition() {
	var x, y, width, height int32
	var rtDesk, rtWindow win.RECT
	win.GetWindowRect(win.GetDesktopWindow(), &rtDesk)
	win.GetWindowRect(wm.GetHWND(), &rtWindow)

	width = rtWindow.Right - rtWindow.Left
	height = rtWindow.Bottom - rtWindow.Top

	x = rtDesk.Right - width
	y = rtDesk.Bottom - (height + 40)

	win.MoveWindow(wm.GetHWND(), x, y, width, height, true)
}

/**
*	center
**/
func (wm *WalkUI) center() {
	//
	var x, y, width, height int32
	var rtDesk, rtWindow win.RECT
	win.GetWindowRect(win.GetDesktopWindow(), &rtDesk)
	win.GetWindowRect(wm.GetHWND(), &rtWindow)

	width = rtWindow.Right - rtWindow.Left
	height = rtWindow.Bottom - rtWindow.Top
	x = (rtDesk.Right - width) / 2
	y = (rtDesk.Bottom - height) / 2

	win.MoveWindow(wm.GetHWND(), x, y, width, height, true)
	//
}

/**
*	Starting
**/
func (wm *WalkUI) Starting(startingFunc WinStartFunc) {
	wm.startingFunc = startingFunc
}

func (wm *WalkUI) Closing(closingFunc WinCloseFunc) {
	wm.closingFunc = closingFunc
}

/**
*	GetHWND
**/
func (wm *WalkUI) GetHWND() win.HWND {
	return wm.window.Handle()
}

/**
*	Window
**/
func (wm *WalkUI) Window() *walk.MainWindow {
	return wm.window
}

/**
*	SetTitle
**/
func (wm *WalkUI) SetTitle(title string) *WalkUI {
	wm.window.SetTitle(title)
	return wm
}

/**
*	SetSize
**/
func (wm *WalkUI) SetSize(width int, height int) *WalkUI {
	wm.window.SetWidth(width)
	wm.window.SetHeight(height)
	return wm
}

/**
*	SetMinSize
**/
func (wm *WalkUI) SetMinSize(width int, height int) *WalkUI {
	maxSize := wm.window.MaxSize()
	wm.window.SetMinMaxSize(walk.Size{Width: width, Height: height}, maxSize)
	return wm
}

/**
*	SetMaxSize
**/
func (wm *WalkUI) SetMaxSize(width int, height int) *WalkUI {
	minSize := wm.window.MinSize()
	wm.window.SetMinMaxSize(minSize, walk.Size{Width: width, Height: height})
	return wm
}

/**
* NoResize
**/
func (wm *WalkUI) NoResize() *WalkUI {
	defStyle := win.GetWindowLong(wm.GetHWND(), win.GWL_STYLE)
	newStyle := defStyle &^ win.WS_THICKFRAME
	win.SetWindowLong(wm.GetHWND(), win.GWL_STYLE, newStyle)
	return wm
}

/**
*	DisableMinBox
**/
func (wm *WalkUI) DisableMinBox() *WalkUI {
	defStyle := win.GetWindowLong(wm.GetHWND(), win.GWL_STYLE)
	newStyle := defStyle &^ win.WS_MINIMIZEBOX
	win.SetWindowLong(wm.GetHWND(), win.GWL_STYLE, newStyle)
	return wm
}

/**
*	DisableMaxBox
**/
func (wm *WalkUI) DisableMaxBox() *WalkUI {
	defStyle := win.GetWindowLong(wm.GetHWND(), win.GWL_STYLE)
	newStyle := defStyle &^ win.WS_MAXIMIZEBOX
	win.SetWindowLong(wm.GetHWND(), win.GWL_STYLE, newStyle)
	return wm
}

/**
*	DisableMinMaxBox
**/
func (wm *WalkUI) DisableMinMaxBox() *WalkUI {
	defStyle := win.GetWindowLong(wm.GetHWND(), win.GWL_STYLE)
	newStyle := defStyle &^ win.WS_MINIMIZEBOX
	newStyle = newStyle &^ win.WS_MAXIMIZEBOX
	win.SetWindowLong(wm.GetHWND(), win.GWL_STYLE, newStyle)
	return wm
}

/**
*	DisableSysmenu
**/
func (wm *WalkUI) DisableSysmenu() *WalkUI {
	defStyle := win.GetWindowLong(wm.GetHWND(), win.GWL_STYLE)
	newStyle := defStyle &^ win.WS_SYSMENU
	win.SetWindowLong(wm.GetHWND(), win.GWL_STYLE, newStyle)
	return wm
}

/**
*	DisableTitleBar
**/
func (wm *WalkUI) DisableTitleBar() *WalkUI {
	defStyle := win.GetWindowLong(wm.GetHWND(), win.GWL_STYLE)
	newStyle := defStyle &^ win.WS_CAPTION
	win.SetWindowLong(wm.GetHWND(), win.GWL_STYLE, newStyle)
	return wm
}

/**
*	SetForeground
**/
func (wm *WalkUI) SetForeground() {
	win.SetForegroundWindow(wm.GetHWND())
}

/**
*	Close
**/
func (wm *WalkUI) Close() {
	wm.Sync(func() {
		wm.window.SetVisible(false)
		wm.window.Close()
	})
}

/**
*	Start
**/
func (wm *WalkUI) Start() {
	wm.window.Show()
	wm.window.Run()
}

/**
*	StartForeground
**/
func (wm *WalkUI) StartForeground() {
	wm.SetForeground()
	wm.window.Show()
	wm.window.Run()
}

/**
*	HideStart
**/
func (wm *WalkUI) HideStart() {
	wm.window.Hide()
	wm.window.Run()
}

/**
*	Hide
**/
func (wm *WalkUI) Hide() {
	wm.window.Hide()
}

/**
*	Show
**/
func (wm *WalkUI) Show() {
	wm.window.Show()
}

/**
*	IgnoreClosing
**/
func (wm *WalkUI) IgnoreClosing() {
	wm.IsIgnoreClose = true
}

/**
*	Sync
**/
func (wm *WalkUI) Sync(syncFunc func()) {
	wm.window.Synchronize(syncFunc)
}

/**
* Append
**/
func (wm *WalkUI) Append(item walk.Widget) {
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
func (wm *WalkUI) Parent() walk.Container {
	if wm.parentList.Len() > 0 {
		parent := wm.parentList.Back().Value.(walk.Container)
		return parent
	} else {
		return wm.window
	}
}

/**
*	End
**/
func (wm *WalkUI) End() {
	if wm.parentList.Len() > 0 {
		popData := wm.parentList.Remove(wm.parentList.Back())
		parent := wm.Parent()
		parent.Children().Add(popData.(walk.Widget))
	}
}

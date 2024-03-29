package walkmgr

import (
	"container/list"

	"github.com/pirogom/walk"
	"github.com/pirogom/win"
)

type WinCloseFunc func() bool
type WinStartFunc func()

var (
	useWalkPositionMgr bool
)

/**
*	WalkUI
**/
type WalkUI struct {
	dialog        *walk.Dialog
	window        *walk.MainWindow
	parentList    *list.List
	IsIgnoreClose bool
	startingFunc  WinStartFunc
	closingFunc   WinCloseFunc
	tb            *walk.ToolBar
}

/**
*	SetUseWalkPositionMgr
**/
func SetUseWalkPositionMgr() {
	useWalkPositionMgr = true
}

/**
*	UseWalkPositionMgr
**/
func UseWalkPositionMgr() bool {
	return useWalkPositionMgr
}

/**
*	CreateWindow
**/
func CreateWindow(title string, posX, posY int, width, height int, margin *walk.Margins, lt ...LayoutType) *WalkUI {
	wm := WalkUI{}

	// set config
	cfg := walk.MainWindowCfg{}
	cfg.Name = defaultWinName
	cfg.Bounds.SetLocation(walk.Point{X: posX, Y: posY})
	cfg.Bounds.SetSize(walk.Size{Width: width, Height: height})

	window, winErr := walk.NewMainWindowWithCfg(&cfg)

	if winErr != nil {
		panic("create window failed. please check manifest and .syso")
	}
	wm.window = window
	wm.dialog = nil
	wm.parentList = list.New()

	wm.window.SetTitle(title)
	wm.window.SetWidth(width)
	wm.window.SetHeight(height)

	var layout walk.Layout

	if len(lt) > 0 {
		// set layout
		switch lt[0] {
		case LAYOUT_VERT:
			layout = walk.NewVBoxLayout()
		case LAYOUT_HORI:
			layout = walk.NewHBoxLayout()
		case LAYOUT_FLOW:
			layout = walk.NewFlowLayout()
		default:
			layout = walk.NewVBoxLayout()
		}
	} else {
		layout = walk.NewVBoxLayout()
	}
	if margin != nil {
		layout.SetMargins(*margin)
	}
	wm.window.SetLayout(layout)

	if defaultIcon != nil {
		wm.window.SetIcon(defaultIcon)
	}

	// windows start
	wm.window.Starting().Attach(func() {
		wm.MoveCenter()
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
*	NewWin
**/
func NewWin(title string, width int, height int, lt ...LayoutType) *WalkUI {
	x, y := CenterPos(width, height)

	return CreateWindow(title, x, y, width, height, nil, lt...)
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
*	NewMargin
**/
func NewMargin(title string, width int, height int, margin walk.Margins, lt ...LayoutType) *WalkUI {
	x, y := CenterPos(width, height)

	return CreateWindow(title, x, y, width, height, &margin, lt...)
}

/**
*	NewAds
**/
func NewAds(title string, width int, height int) *WalkUI {
	wm := WalkUI{}

	// set config
	cfg := walk.MainWindowCfg{}
	cfg.Name = defaultWinName
	posX, posY := AdsPos(width, height)
	cfg.Bounds.SetLocation(walk.Point{X: posX, Y: posY})
	cfg.Bounds.SetSize(walk.Size{Width: width, Height: height})

	window, winErr := walk.NewMainWindowWithCfg(&cfg)

	if winErr != nil {
		panic("create window failed. please check manifest and .syso")
	}

	window.DisablePositionMgr()

	wm.window = window
	wm.dialog = nil
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
		wm.MoveAds()
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
*	WindowPos
**/
func (wm *WalkUI) WindowPos() (int, int, int, int) {
	var rtWindow win.RECT

	win.GetWindowRect(wm.GetHWND(), &rtWindow)

	width := int(rtWindow.Right - rtWindow.Left)
	height := int(rtWindow.Bottom - rtWindow.Top)

	x := int(rtWindow.Left)
	y := int(rtWindow.Top)

	return x, y, width, height
}

/**
*	CenterPos
**/
func CenterPos(width int, height int) (int, int) {
	if UseWalkPositionMgr() && walk.PosMgr.HasPosition() {
		pmX, pmY, pmW, pmH, _, deskH := walk.PosMgr.Get()

		var rX, rY int

		if pmW > width {
			rX = pmX + ((pmW - width) / 2)
		} else if pmW < width {
			rX = pmX - ((width - pmW) / 2)
		} else {
			rX = pmX
		}
		if pmH > height {
			rY = pmY + ((pmH - height) / 2)
		} else if pmH < height {
			rY = pmY - ((height - pmH) / 2)
		} else {
			rY = pmY
		}

		if rX < 0 {
			rX = 0
		}
		if rY < 0 {
			rY = 0
		} else if rY+height > deskH {
			rY = deskH - height
		}
		return rX, rY
	}
	var x, y int
	var rtDesk win.RECT
	win.GetWindowRect(win.GetDesktopWindow(), &rtDesk)

	x = (int(rtDesk.Right) - width) / 2
	y = (int(rtDesk.Bottom) - height) / 2

	return x, y
}

/**
*	AdsPos
**/
func AdsPos(width int, height int) (int, int) {
	var x, y int
	var rtDesk win.RECT
	win.GetWindowRect(win.GetDesktopWindow(), &rtDesk)

	x = int(rtDesk.Right) - width
	y = int(rtDesk.Bottom) - (height + 40)

	return x, y
}

/**
*	MoveAds
**/
func (wm *WalkUI) MoveAds() {
	var rtWindow win.RECT
	win.GetWindowRect(wm.GetHWND(), &rtWindow)

	width := rtWindow.Right - rtWindow.Left
	height := rtWindow.Bottom - rtWindow.Top

	x, y := AdsPos(int(width), int(height))

	win.MoveWindow(wm.GetHWND(), int32(x), int32(y), width, height, true)
}

/**
*	MoveCenter
**/
func (wm *WalkUI) MoveCenter() {
	var rtWindow win.RECT
	win.GetWindowRect(wm.GetHWND(), &rtWindow)

	width := rtWindow.Right - rtWindow.Left
	height := rtWindow.Bottom - rtWindow.Top

	x, y := CenterPos(int(width), int(height))

	win.MoveWindow(wm.GetHWND(), int32(x), int32(y), width, height, true)
	//
}

/**
*	Starting
**/
func (wm *WalkUI) Starting(startingFunc WinStartFunc) {
	wm.startingFunc = startingFunc
}

/**
*	Closing
**/
func (wm *WalkUI) Closing(closingFunc WinCloseFunc) {
	wm.closingFunc = closingFunc
}

/**
*	GetHWND
**/
func (wm *WalkUI) GetHWND() win.HWND {
	if wm.dialog != nil {
		return wm.dialog.Handle()
	}
	return wm.window.Handle()
}

/**
*	Window
**/
func (wm *WalkUI) Window() *walk.MainWindow {
	if wm.dialog != nil {
		panic("Window() method is not support dialog window. Use Dlg() method")
		return nil
	}
	return wm.window
}

/**
*	SetTitle
**/
func (wm *WalkUI) SetTitle(title string) *WalkUI {
	if wm.dialog != nil {
		wm.dialog.SetTitle(title)
		return wm
	}
	wm.window.SetTitle(title)
	return wm
}

/**
*	SetSize
**/
func (wm *WalkUI) SetSize(width int, height int) *WalkUI {
	if wm.dialog != nil {
		wm.dialog.SetWidth(width)
		wm.dialog.SetHeight(height)
		return wm
	}
	wm.window.SetWidth(width)
	wm.window.SetHeight(height)
	return wm
}

/**
*	SetMinSize
**/
func (wm *WalkUI) SetMinSize(width int, height int) *WalkUI {
	maxSize := wm.window.MaxSize()
	if wm.dialog != nil {
		wm.dialog.SetMinMaxSize(walk.Size{Width: width, Height: height}, maxSize)
		return wm
	}
	wm.window.SetMinMaxSize(walk.Size{Width: width, Height: height}, maxSize)
	return wm
}

/**
*	SetMaxSize
**/
func (wm *WalkUI) SetMaxSize(width int, height int) *WalkUI {
	minSize := wm.window.MinSize()
	if wm.dialog != nil {
		wm.dialog.SetMinMaxSize(minSize, walk.Size{Width: width, Height: height})
		return wm
	}
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
	if wm.dialog != nil {
		panic("Close() method is not support dialog window. Use CloseDLG() method.")
		return
	}
	wm.Sync(func() {
		wm.window.SetVisible(false)
		wm.window.Close()
	})
}

/**
*	Start
**/
func (wm *WalkUI) Start() {
	if wm.dialog != nil {
		panic("Start() method is not support dialog window. Use StartDLG() method")
		return
	}
	wm.window.Show()
	wm.window.Run()
}

/**
*	StartForeground
**/
func (wm *WalkUI) StartForeground() {
	if wm.dialog != nil {
		panic("StartForeground() method is not support dialog window. Use StartForegroundDLG() method")
		return
	}
	wm.SetForeground()
	wm.window.Show()
	wm.window.Run()
}

/**
*	HideStart
**/
func (wm *WalkUI) HideStart() {
	if wm.dialog != nil {
		panic("HideStart() method is not support dialog window. Use HideStartDLG() method")
		return
	}
	wm.window.Hide()
	wm.window.Run()
}

/**
*	Hide
**/
func (wm *WalkUI) Hide() {
	if wm.dialog != nil {
		wm.dialog.Hide()
		return
	}
	wm.window.Hide()
}

/**
*	Show
**/
func (wm *WalkUI) Show() {
	if wm.dialog != nil {
		wm.dialog.Show()
		return
	}
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
	if wm.dialog != nil {
		wm.dialog.Synchronize(syncFunc)
		return
	}
	wm.window.Synchronize(syncFunc)
}

/**
* Append
**/
func (wm *WalkUI) Append(item walk.Widget) {
	if wm.parentList.Len() == 0 {
		if wm.dialog != nil {
			wm.dialog.Children().Add(item)
		} else {
			wm.window.Children().Add(item)
		}
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
		if wm.dialog != nil {
			return wm.dialog
		} else {
			return wm.window
		}
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

/**
*	AddMenu
**/
func (wm *WalkUI) AddMenu(in *MenuMgr) {
	if wm.dialog != nil {
		panic("dialog window does not support menu.")
		return
	}
	wm.window.Menu().Actions().Add(in.MenuAct)
}



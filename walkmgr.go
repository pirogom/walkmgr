package walkmgr

import (
	"container/list"

	"github.com/pirogom/walk"
	"github.com/pirogom/win"
)

/**
*	walk_ui
**/
type walk_ui struct {
	window       *walk.MainWindow
	parentList   *list.List
	startingFunc func()
}

/**
*	NewWin
**/
func NewWin(title string, width int, height int, lt ...LayoutType) *walk_ui {
	wm := walk_ui{}

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
		default:
			wm.window.SetLayout(walk.NewVBoxLayout())
		}
	} else {
		wm.window.SetLayout(walk.NewVBoxLayout())
	}

	if defaultIcon != nil {
		wm.window.SetIcon(defaultIcon)
	}

	wm.window.Starting().Attach(func() {
		wm.center()
		if wm.startingFunc != nil {
			wm.startingFunc()
		}
	})

	return &wm
}

/**
*	NewAds
**/
func NewAds(title string, width int, height int) *walk_ui {
	wm := walk_ui{}

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

	wm.window.Starting().Attach(func() {
		wm.adsPosition()

		if wm.startingFunc != nil {
			wm.startingFunc()
		}
	})

	return &wm
}

/**
*	adsPosition
**/
func (wm *walk_ui) adsPosition() {
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
func (wm *walk_ui) center() {
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
func (wm *walk_ui) Starting(startingFunc func()) {
	wm.startingFunc = startingFunc
}

/**
*	GetHWND
**/
func (wm *walk_ui) GetHWND() win.HWND {
	return wm.window.Handle()
}

/**
*	Window
**/
func (wm *walk_ui) Window() *walk.MainWindow {
	return wm.window
}

/**
*	SetTitle
**/
func (wm *walk_ui) SetTitle(title string) *walk_ui {
	wm.window.SetTitle(title)
	return wm
}

/**
*	SetSize
**/
func (wm *walk_ui) SetSize(width int, height int) *walk_ui {
	wm.window.SetWidth(width)
	wm.window.SetHeight(height)
	return wm
}

/**
*	SetMinSize
**/
func (wm *walk_ui) SetMinSize(width int, height int) *walk_ui {
	maxSize := wm.window.MaxSize()
	wm.window.SetMinMaxSize(walk.Size{Width: width, Height: height}, maxSize)
	return wm
}

/**
*	SetMaxSize
**/
func (wm *walk_ui) SetMaxSize(width int, height int) *walk_ui {
	minSize := wm.window.MinSize()
	wm.window.SetMinMaxSize(minSize, walk.Size{Width: width, Height: height})
	return wm
}

/**
* NoResize
**/
func (wm *walk_ui) NoResize() *walk_ui {
	defStyle := win.GetWindowLong(wm.GetHWND(), win.GWL_STYLE)
	newStyle := defStyle &^ win.WS_THICKFRAME
	win.SetWindowLong(wm.GetHWND(), win.GWL_STYLE, newStyle)
	return wm
}

/**
*	DisableMinBox
**/
func (wm *walk_ui) DisableMinBox() *walk_ui {
	defStyle := win.GetWindowLong(wm.GetHWND(), win.GWL_STYLE)
	newStyle := defStyle &^ win.WS_MINIMIZEBOX
	win.SetWindowLong(wm.GetHWND(), win.GWL_STYLE, newStyle)
	return wm
}

/**
*	DisableMaxBox
**/
func (wm *walk_ui) DisableMaxBox() *walk_ui {
	defStyle := win.GetWindowLong(wm.GetHWND(), win.GWL_STYLE)
	newStyle := defStyle &^ win.WS_MAXIMIZEBOX
	win.SetWindowLong(wm.GetHWND(), win.GWL_STYLE, newStyle)
	return wm
}

/**
*	DisableMinMaxBox
**/
func (wm *walk_ui) DisableMinMaxBox() *walk_ui {
	defStyle := win.GetWindowLong(wm.GetHWND(), win.GWL_STYLE)
	newStyle := defStyle &^ win.WS_MINIMIZEBOX
	newStyle = newStyle &^ win.WS_MAXIMIZEBOX
	win.SetWindowLong(wm.GetHWND(), win.GWL_STYLE, newStyle)
	return wm
}

/**
*	DisableSysmenu
**/
func (wm *walk_ui) DisableSysmenu() *walk_ui {
	defStyle := win.GetWindowLong(wm.GetHWND(), win.GWL_STYLE)
	newStyle := defStyle &^ win.WS_SYSMENU
	win.SetWindowLong(wm.GetHWND(), win.GWL_STYLE, newStyle)
	return wm
}

/**
*	DisableTitleBar
**/
func (wm *walk_ui) DisableTitleBar() *walk_ui {
	defStyle := win.GetWindowLong(wm.GetHWND(), win.GWL_STYLE)
	newStyle := defStyle &^ win.WS_CAPTION
	win.SetWindowLong(wm.GetHWND(), win.GWL_STYLE, newStyle)
	return wm
}

/**
*	SetForeground
**/
func (wm *walk_ui) SetForeground() {
	win.SetForegroundWindow(wm.GetHWND())
}

/**
*	Close
**/
func (wm *walk_ui) Close() {
	wm.window.Close()
}

/**
*	Start
**/
func (wm *walk_ui) Start() {
	wm.window.Show()
	wm.window.Run()
}

/**
*	StartForeground
**/
func (wm *walk_ui) StartForeground() {
	wm.SetForeground()
	wm.window.Show()
	wm.window.Run()
}

/**
*	HideStart
**/
func (wm *walk_ui) HideStart() {
	wm.window.Hide()
	wm.window.Run()
}

/**
*	Hide
**/
func (wm *walk_ui) Hide() {
	wm.window.Hide()
}

/**
*	Show
**/
func (wm *walk_ui) Show() {
	wm.window.Show()
}

/**
*	IgnoreClosing
**/
func (wm *walk_ui) IgnoreClosing() {
	wm.window.Closing().Attach(func(canceled *bool, reason walk.CloseReason) {
		if wm.window.Visible() {
			*canceled = true
		}
	})
}

/**
*	Sync
**/
func (wm *walk_ui) Sync(syncFunc func()) {
	wm.window.Synchronize(syncFunc)
}

/**
*	ForceClose
**/
func (wm *walk_ui) ForceClose() {
	wm.Sync(func() {
		wm.window.SetVisible(false)
		wm.window.Close()
	})
}

/**
* Append
**/
func (wm *walk_ui) Append(item walk.Widget) {
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
func (wm *walk_ui) Parent() walk.Container {
	if wm.parentList.Len() > 0 {
		parent := wm.parentList.Back().Value.(walk.Container)
		return parent
	} else {
		return wm.window
	}
}

/**
*	HSplit
**/
func (wm *walk_ui) HSplit() *walk.Splitter {
	return wm.Split(LAYOUT_HORI)
}

/**
*	VSplit
**/
func (wm *walk_ui) VSplit() *walk.Splitter {
	return wm.Split(LAYOUT_VERT)
}

/**
*	Split
**/
func (wm *walk_ui) Split(lt ...LayoutType) *walk.Splitter {
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
	return hs
}

/**
*	GroupBox
**/
func (wm *walk_ui) GroupBox(title string, lt ...LayoutType) (*walk_ui, *walk.GroupBox) {
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
func (wm *walk_ui) End() {
	if wm.parentList.Len() > 0 {
		popData := wm.parentList.Remove(wm.parentList.Back())
		parent := wm.Parent()
		parent.Children().Add(popData.(walk.Widget))
	}
}

/**
*	EndSplit
**/
func (wm *walk_ui) EndSplit() {
	wm.End()
}

/**
*	EndGroupBox
**/
func (wm *walk_ui) EndGroupBox() {
	wm.End()
}

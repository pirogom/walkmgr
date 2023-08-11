package walkmgr

import (
	"container/list"

	"github.com/pirogom/walk"
)

/**
*	NewDialog
**/
func NewDialog(owner *walk.MainWindow, title string, width, height int, margin *walk.Margins, lt ...LayoutType) *WalkUI {
	wm := WalkUI{}

	dlg, dlgErr := walk.NewDialog(owner)

	if dlgErr != nil {
		panic(dlgErr.Error())
	}

	wm.dialog = dlg
	wm.window = nil
	wm.parentList = list.New()

	wm.dialog.SetTitle(title)
	wm.dialog.SetWidth(width)
	wm.dialog.SetHeight(height)
	wm.dialog.SetMinMaxSize(walk.Size{Width: width, Height: height}, walk.Size{Width: width, Height: height})
	wm.NoResize()

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
	wm.dialog.SetLayout(layout)

	if defaultIcon != nil {
		wm.dialog.SetIcon(defaultIcon)
	}

	// windows start
	wm.dialog.Starting().Attach(func() {
		if wm.startingFunc != nil {
			wm.startingFunc()
		}
	})

	// closing
	wm.dialog.Closing().Attach(func(canceled *bool, reason walk.CloseReason) {
		if wm.IsIgnoreClose && wm.Dlg().Visible() {
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
*	GetDlg
**/
func (wm *WalkUI) Dlg() *walk.Dialog {
	if wm.dialog == nil {
		panic("Dlg() method is not support non dialog window. Use Window() method.")
	}
	return wm.dialog
}

/**
*	CloseDLG
**/
func (wm *WalkUI) CloseDLG(res int) {
	if wm.dialog == nil {
		panic("CloseDLG method is not support non dialog window. Use Close() method")
	}
	wm.Sync(func() {
		wm.dialog.SetVisible(false)
		wm.dialog.Close(res)
	})
}

/**
*	StartDLG
**/
func (wm *WalkUI) StartDLG() int {
	if wm.dialog == nil {
		panic("StartDLG method is not support non dialog window. Use Start() method")
	}
	wm.dialog.Show()
	return wm.dialog.Run()
}

/**
*	StartForegroundDLG
**/
func (wm *WalkUI) StartForegroundDLG() int {
	if wm.dialog == nil {
		panic("StartForegroundDLG method is not support non dialog window. Use StartForeground() method")
	}
	wm.SetForeground()
	wm.dialog.Show()
	return wm.dialog.Run()
}

/**
*	HideStartDLG
**/
func (wm *WalkUI) HideStartDLG() int {
	if wm.dialog == nil {
		panic("HideStartDLG method is not support non dialog window. Use HideStart() method")
	}
	wm.dialog.Hide()
	return wm.dialog.Run()
}

/**
*	DlgResult
**/
func (wm *WalkUI) DlgResult() int {
	if wm.dialog == nil {
		panic("DlgResult() method is not support non dialog window.")
	}
	return wm.dialog.Result()
}

/**
*	SetOkButton
**/
func (wm *WalkUI) SetOkButton(btn *walk.PushButton) error {
	if wm.dialog == nil {
		panic("SetOkButton does not support non dialog window.")
	}
	return wm.dialog.SetDefaultButton(btn)
}

/**
*	SetCancelButton
**/
func (wm *WalkUI) SetCancelButton(btn *walk.PushButton) error {
	if wm.dialog == nil {
		panic("SetCancelButton does not support non dialog window.")
	}
	return wm.dialog.SetCancelButton(btn)
}

func (wm *WalkUI) Accept() {
	if wm.dialog == nil {
		panic("do not use for non dialog window")
	}
	wm.dialog.Accept()
}

func (wm *WalkUI) Cancel() {
	if wm.dialog == nil {
		panic("do not use for non dialog window")
	}
	wm.dialog.Cancel()
}

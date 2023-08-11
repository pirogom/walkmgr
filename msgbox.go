package walkmgr

import (
	"github.com/pirogom/walk"
	"github.com/pirogom/win"
)

/**
 *	MsgBox
 */
func (wm *WalkUI) MsgBox(msg string) {
	if wm.dialog != nil {
		walk.MsgBox(wm.dialog, "알림", msg, walk.MsgBoxOK|walk.MsgBoxSetForeground)
	} else {
		walk.MsgBox(wm.window, "알림", msg, walk.MsgBoxOK|walk.MsgBoxSetForeground)
	}
}

/**
 * 	Confirm
 */
func (wm *WalkUI) Confirm(msg string) bool {
	if wm.dialog != nil {
		if walk.MsgBox(wm.dialog, "알림", msg, walk.MsgBoxYesNo|walk.MsgBoxSetForeground) == win.IDNO {
			return false
		} else {
			return true
		}
	} else {
		if walk.MsgBox(wm.window, "알림", msg, walk.MsgBoxYesNo|walk.MsgBoxSetForeground) == win.IDNO {
			return false
		} else {
			return true
		}
	}
}

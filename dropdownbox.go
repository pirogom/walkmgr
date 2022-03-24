package walkmgr

import "github.com/pirogom/walk"

/**
*	DropDownBox
**/
func (wm *WalkUI) DropDownBox(data []string, defIdx ...int) *walk.ComboBox {
	cb, _ := walk.NewDropDownBox(wm.Parent())
	cb.SetModel(data)

	if len(defIdx) == 0 {
		cb.SetCurrentIndex(0)
	} else {
		cb.SetCurrentIndex(defIdx[0])
	}

	wm.Append(cb)
	return cb
}

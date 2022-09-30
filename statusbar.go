package walkmgr

import "github.com/pirogom/walk"

/**
*	StatusBar
**/
func (wm *WalkUI) StatusBar() *walk.StatusBar {
	nb, _ := walk.NewStatusBar(wm.Parent())

	wm.Append(nb)

	return nb
}

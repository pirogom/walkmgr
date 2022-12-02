package walkmgr

import "github.com/pirogom/walk"

/**
*	StatusBarItems
**/
func (wm *WalkUI) StatusBarItems() *walk.StatusBarItemList {
	return wm.Window().StatusBar().Items()
}

/**
*	StatusBarVisible
**/
func (wm *WalkUI) StatusBarVisible(visible bool) {
	wm.Window().StatusBar().SetVisible(visible)
}

/**
*	AddStatusItem
**/
func (wm *WalkUI) AddStatusItem(text string, width int, icon *walk.Icon, clickedFn func()) *walk.StatusBarItem {
	nb := walk.NewStatusBarItem()
	nb.SetText(text)
	if width > 0 {
		nb.SetWidth(width)
	}
	if icon != nil {
		nb.SetIcon(icon)
	}
	if clickedFn != nil {
		nb.Clicked().Attach(clickedFn)
	}
	wm.Window().StatusBar().Items().Add(nb)
	return nb
}

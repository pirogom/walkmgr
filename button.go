package walkmgr

import "github.com/pirogom/walk"

/**
*	PushButton
**/
func (wm *WalkUI) PushButton(text string, clickFunc func()) *walk.PushButton {
	btn, _ := walk.NewPushButton(wm.Parent())
	btn.SetText(text)
	btn.Clicked().Attach(clickFunc)

	wm.Append(btn)
	return btn
}

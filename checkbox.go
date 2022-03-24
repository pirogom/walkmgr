package walkmgr

import "github.com/pirogom/walk"

/**
*	CheckBox
**/
func (wm *walkmgr) CheckBox(text string, checked bool, attachFunc func()) *walk.CheckBox {
	cb, _ := walk.NewCheckBox(wm.Parent())
	cb.SetText(text)
	cb.SetChecked(checked)
	cb.CheckStateChanged().Attach(attachFunc)

	wm.Append(cb)
	return cb
}

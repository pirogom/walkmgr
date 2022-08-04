package walkmgr

import "github.com/pirogom/walk"

/**
*	RadioButton
**/
func (wm *WalkUI) RadioButton(text string, value interface{}) *walk.RadioButton {
	rdo, _ := walk.NewRadioButton(wm.Parent())
	rdo.SetText(text)
	rdo.SetValue(value)

	wm.Append(rdo)
	return rdo
}

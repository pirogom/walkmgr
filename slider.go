package walkmgr

import "github.com/pirogom/walk"

/**
*	Slider
**/
func (wm *walkmgr) Slider(minVal int, maxVal int, defVal int) *walk.Slider {
	ne, _ := walk.NewSlider(wm.Parent())
	ne.SetRange(minVal, maxVal)
	ne.SetValue(defVal)

	wm.Append(ne)
	return ne
}

package walkmgr

import "github.com/pirogom/walk"

/**
*	GroupBox
**/
func (wm *WalkUI) GroupBox(title string, lt ...LayoutType) *walk.GroupBox {
	gb, _ := walk.NewGroupBox(wm.Parent())

	if title != "" {
		gb.SetTitle(title)
	}

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
	return gb
}

/**
*	HGroupBox
**/
func (wm *WalkUI) HGroupBox(title string) *walk.GroupBox {
	return wm.GroupBox(title, LAYOUT_HORI)
}

/**
*	VGroupBox
**/
func (wm *WalkUI) VGroupBox(title string) *walk.GroupBox {
	return wm.GroupBox(title, LAYOUT_VERT)
}

/**
*	HBox
**/
func (wm *WalkUI) HBox() *walk.GroupBox {
	return wm.GroupBox("", LAYOUT_HORI)
}

/**
*	VBox
**/
func (wm *WalkUI) VBox() *walk.GroupBox {
	return wm.GroupBox("", LAYOUT_VERT)
}

/**
*	EndGroupBox
**/
func (wm *WalkUI) EndGroupBox() {
	wm.End()
}

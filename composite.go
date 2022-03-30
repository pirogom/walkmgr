package walkmgr

import "github.com/pirogom/walk"

/**
*	Composite
**/
func (wm *WalkUI) Composite(lt ...LayoutType) *walk.Composite {
	cs, _ := walk.NewComposite(wm.Parent())

	if len(lt) == 0 {
		cs.SetLayout(walk.NewVBoxLayout())
	} else {
		switch lt[0] {
		case LAYOUT_VERT:
			cs.SetLayout(walk.NewVBoxLayout())
		case LAYOUT_HORI:
			cs.SetLayout(walk.NewHBoxLayout())
		case LAYOUT_FLOW:
			cs.SetLayout(walk.NewFlowLayout())
		default:
			cs.SetLayout(walk.NewVBoxLayout())
		}
	}
	wm.parentList.PushBack(cs)
	return cs
}

/**
*	HComposite
**/
func (wm *WalkUI) HComposite() *walk.Composite {
	return wm.Composite(LAYOUT_HORI)
}

/**
*	VComposite
**/
func (wm *WalkUI) VComposite() *walk.Composite {
	return wm.Composite(LAYOUT_VERT)
}

/**
*	FComposite
**/
func (wm *WalkUI) FComposite() *walk.Composite {
	return wm.Composite(LAYOUT_FLOW)
}

/**
*	EndComposite
**/
func (wm *WalkUI) EndComposite() {
	wm.End()
}

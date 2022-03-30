package walkmgr

import (
	"github.com/pirogom/walk"
)

/**
*	Composite
**/
func (wm *WalkUI) Composite(lt ...LayoutType) *walk.Composite {
	cs, _ := walk.NewComposite(wm.Parent())

	if len(lt) == 0 {
		ly := walk.NewVBoxLayout()
		ly.SetMargins(walk.Margins{0, 0, 0, 0})
		cs.SetLayout(ly)
	} else {
		switch lt[0] {
		case LAYOUT_VERT:
			ly := walk.NewVBoxLayout()
			ly.SetMargins(walk.Margins{0, 0, 0, 0})
			cs.SetLayout(ly)
		case LAYOUT_HORI:
			ly := walk.NewHBoxLayout()
			ly.SetMargins(walk.Margins{0, 0, 0, 0})
			cs.SetLayout(ly)
		case LAYOUT_FLOW:
			ly := walk.NewFlowLayout()
			ly.SetMargins(walk.Margins{0, 0, 0, 0})
			cs.SetLayout(ly)
		default:
			ly := walk.NewVBoxLayout()
			ly.SetMargins(walk.Margins{0, 0, 0, 0})
			cs.SetLayout(ly)
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

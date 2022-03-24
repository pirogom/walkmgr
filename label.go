package walkmgr

import (
	"github.com/pirogom/walk"
	"github.com/pirogom/win"
)

/**
*	MultiLineLabel
**/
func (wm *walk_ui) MultiLineLabel(text string) *walk.Label {
	ne, _ := walk.NewLabelWithStyle(wm.Parent(), win.SS_EDITCONTROL)
	ne.SetText(text)
	ne.SetAlignment(walk.AlignHCenterVCenter)
	ne.SetTextAlignment(walk.AlignCenter)

	wm.Append(ne)
	return ne
}

/**
*	Label
**/
func (wm *walk_ui) Label(text string, at ...AlignType) *walk.Label {
	ne, _ := walk.NewLabel(wm.Parent())

	ne.SetText(text)

	if len(at) == 0 {
		ne.SetTextAlignment(walk.AlignDefault)
	} else {
		switch at[0] {
		case ALIGN_LEFT:
			ne.SetTextAlignment(walk.AlignNear)
		case ALIGN_CENTER:
			ne.SetTextAlignment(walk.AlignCenter)
		case ALIGN_RIGHT:
			ne.SetTextAlignment(walk.AlignFar)
		default:
			ne.SetTextAlignment(walk.AlignDefault)
		}
	}

	wm.Append(ne)
	return ne
}

func (wm *walk_ui) LabelCenter(text string) *walk.Label {
	return wm.Label(text, ALIGN_CENTER)
}

func (wm *walk_ui) LabelRight(text string) *walk.Label {
	return wm.Label(text, ALIGN_RIGHT)
}

func (wm *walk_ui) LabelLeft(text string) *walk.Label {
	return wm.Label(text, ALIGN_LEFT)
}

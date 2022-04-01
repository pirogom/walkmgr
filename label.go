package walkmgr

import (
	"github.com/pirogom/walk"
	"github.com/pirogom/win"
)

/**
*	MultiLineLabel
**/
func (wm *WalkUI) MultiLineLabel(text string, at ...AlignType) *walk.Label {
	ne, _ := walk.NewLabelWithStyle(wm.Parent(), win.SS_EDITCONTROL)
	ne.SetText(text)
	ne.SetAlignment(walk.AlignHCenterVCenter)

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

/**
*	Label
**/
func (wm *WalkUI) Label(text string, at ...AlignType) *walk.Label {
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

func (wm *WalkUI) LabelCenter(text string) *walk.Label {
	return wm.Label(text, ALIGN_CENTER)
}

func (wm *WalkUI) LabelRight(text string) *walk.Label {
	return wm.Label(text, ALIGN_RIGHT)
}

func (wm *WalkUI) LabelLeft(text string) *walk.Label {
	return wm.Label(text, ALIGN_LEFT)
}

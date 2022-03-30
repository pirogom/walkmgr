package walkmgr

import "github.com/pirogom/walk"

/**
*	HSplit
**/
func (wm *WalkUI) HSplit() *walk.Splitter {
	return wm.Split(LAYOUT_HORI)
}

/**
*	VSplit
**/
func (wm *WalkUI) VSplit() *walk.Splitter {
	return wm.Split(LAYOUT_VERT)
}

/**
*	Split
**/
func (wm *WalkUI) Split(lt ...LayoutType) *walk.Splitter {
	var hs *walk.Splitter

	if len(lt) == 0 {
		hs, _ = walk.NewHSplitter(wm.Parent())
	} else {
		switch lt[0] {
		case LAYOUT_VERT:
			hs, _ = walk.NewVSplitter(wm.Parent())
		case LAYOUT_HORI:
			hs, _ = walk.NewHSplitter(wm.Parent())
		default:
			hs, _ = walk.NewHSplitter(wm.Parent())
		}
	}
	wm.parentList.PushBack(hs)
	return hs
}

/**
*	EndSplit
**/
func (wm *WalkUI) EndSplit() {
	wm.End()
}

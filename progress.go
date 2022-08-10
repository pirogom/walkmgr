package walkmgr

import "github.com/pirogom/walk"

/**
*	ProgressBar
**/
func (wm *WalkUI) ProgressBar(min, max, curr int) *walk.ProgressBar {
	pb, _ := walk.NewProgressBar(wm.Parent())

	pb.SetValue(curr)
	pb.SetRange(min, max)
	wm.Append(pb)

	return pb
}

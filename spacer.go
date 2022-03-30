package walkmgr

import "github.com/pirogom/walk"

/**
*	HSpacer
**/
func (wm *WalkUI) HSpacer() *walk.Spacer {
	hs, _ := walk.NewHSpacer(wm.Parent())
	wm.Append(hs)
	return hs
}

/**
*	HSpacerFixed
**/
func (wm *WalkUI) HSpacerFixed(width int) *walk.Spacer {
	hs, _ := walk.NewHSpacerFixed(wm.Parent(), width)
	wm.Append(hs)
	return hs
}

/**
*	VSpacer
**/
func (wm *WalkUI) VSpacer() *walk.Spacer {
	vs, _ := walk.NewVSpacer(wm.Parent())
	wm.Append(vs)
	return vs
}

/**
*	HSpacerFixed
**/
func (wm *WalkUI) VSpacerFixed(height int) *walk.Spacer {
	vs, _ := walk.NewVSpacerFixed(wm.Parent(), height)
	wm.Append(vs)
	return vs
}

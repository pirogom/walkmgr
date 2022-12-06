package walkmgr

import "github.com/pirogom/walk"

/**
*	ScrollView
**/
func (wm *WalkUI) ScrollView(hori bool, vert bool, lt ...LayoutType) *walk.ScrollView {
	sv, _ := walk.NewScrollView(wm.Parent())
	sv.SetScrollbars(hori, vert)
	if len(lt) == 0 {
		ly := walk.NewVBoxLayout()
		ly.SetMargins(walk.Margins{0, 0, 0, 0})
		sv.SetLayout(ly)
	} else {
		switch lt[0] {
		case LAYOUT_VERT:
			ly := walk.NewVBoxLayout()
			ly.SetMargins(walk.Margins{0, 0, 0, 0})
			sv.SetLayout(ly)
		case LAYOUT_HORI:
			ly := walk.NewHBoxLayout()
			ly.SetMargins(walk.Margins{0, 0, 0, 0})
			sv.SetLayout(ly)
		case LAYOUT_FLOW:
			ly := walk.NewFlowLayout()
			ly.SetMargins(walk.Margins{0, 0, 0, 0})
			sv.SetLayout(ly)
		default:
			ly := walk.NewVBoxLayout()
			ly.SetMargins(walk.Margins{0, 0, 0, 0})
			sv.SetLayout(ly)
		}
	}
	wm.parentList.PushBack(sv)
	return sv
}

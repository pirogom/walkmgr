package walkmgr

import (
	"os/exec"

	"github.com/pirogom/walk"
)

/**
*	LinkLabel
**/
func (wm *WalkUI) LinkLabel(text string, at ...AlignType) *walk.LinkLabel {
	ne, _ := walk.NewLinkLabel(wm.Parent())

	ne.SetText(text)

	if len(at) == 0 {
		ne.SetAlignment(walk.AlignHVDefault)
	} else {
		switch at[0] {
		case ALIGN_LEFT:
			ne.SetAlignment(walk.AlignHNearVCenter)
		case ALIGN_CENTER:
			ne.SetAlignment(walk.AlignHCenterVCenter)
		case ALIGN_RIGHT:
			ne.SetAlignment(walk.AlignHFarVCenter)
		default:
			ne.SetAlignment(walk.AlignHVDefault)
		}
	}
	wm.Append(ne)

	ne.LinkActivated().Attach(func(link *walk.LinkLabelLink) {
		exec.Command("rundll32.exe", "url.dll,FileProtocolHandler", link.URL()).Start()
	})

	return ne
}

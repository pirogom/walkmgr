package walkmgr

import "github.com/pirogom/walk"

type LayoutType int8
type AlignType int8

const (
	LAYOUT_VERT LayoutType = 0 // Vertical
	LAYOUT_HORI            = 1 // Horizontal
)

const (
	ALIGN_LEFT   AlignType = 0
	ALIGN_CENTER           = 1
	ALIGN_RIGHT            = 2
)

const (
	IV_IDEAL   walk.ImageViewMode = walk.ImageViewModeIdeal
	IV_CORNER                     = walk.ImageViewModeCorner
	IV_CENTER                     = walk.ImageViewModeCenter
	IV_SHRINK                     = walk.ImageViewModeShrink
	IV_ZOOM                       = walk.ImageViewModeZoom
	IV_STRETCH                    = walk.ImageViewModeStretch
)

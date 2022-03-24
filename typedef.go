package walkmgr

type LayoutType int8
type AlignType int8

const (
	LAYOUT_VERT LayoutType = 0 // Vertical
	LAYOUT_HORI LayoutType = 1 // Horizontal

	ALIGN_LEFT   AlignType = 0
	ALIGN_CENTER AlignType = 1
	ALIGN_RIGHT  AlignType = 2
)

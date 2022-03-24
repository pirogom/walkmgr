package walkmgr

import "github.com/pirogom/walk"

/**
*	ImageView
**/
func (wm *WalkUI) ImageView(mode ...walk.ImageViewMode) *walk.ImageView {
	iv, _ := walk.NewImageView(wm.Parent())

	if len(mode) > 0 {
		iv.SetMode(mode[0])
	}

	wm.Append(iv)
	return iv
}

/**
*	LoadImage
**/
func LoadImage(fileName string) *walk.Image {
	retImage, retImageErr := walk.NewImageFromFile(fileName)

	if retImageErr != nil {
		return nil
	}
	return &retImage
}

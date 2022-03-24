package walkmgr

import "github.com/pirogom/walk"

/**
*	ImageView
**/
func (wm *walkmgr) ImageView() *walk.ImageView {
	iv, _ := walk.NewImageView(wm.Parent())
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

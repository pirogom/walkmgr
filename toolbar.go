package walkmgr

import "github.com/pirogom/walk"

/**
*	NewImageToolBar
**/
func (wm *WalkUI) NewToolBarWithStyle(style walk.ToolBarButtonStyle) error {
	tb, err := walk.NewToolBarWithOrientationAndButtonStyle(wm.Window(), walk.Horizontal, style)

	if err != nil {
		return err
	}

	wm.tb = tb
	return nil
}

/**
*	NewToolBar
**/
func (wm *WalkUI) NewToolBar() error {
	return wm.NewToolBarWithStyle(walk.ToolBarButtonImageOnly)
}

/**
*	NewToolBarWithText
**/
func (wm *WalkUI) NewToolBarWithText() error {
	return wm.NewToolBarWithStyle(walk.ToolBarButtonImageBeforeText)
}

/**
*	NewToolBarWithText2
**/
func (wm *WalkUI) NewToolBarWithText2() error {
	return wm.NewToolBarWithStyle(walk.ToolBarButtonImageAboveText)
}

/**
*	ToolBar
**/
func (wm *WalkUI) ToolBar() *walk.ToolBar {
	return wm.tb
}

/**
*	AddTool
**/
func (wm *WalkUI) AddTool(title string, icoPath string, trigFunc func()) *walk.Action {
	if wm.tb == nil {
		panic("please InitToolBar first!")
	}
	act := walk.NewAction()
	act.SetText(title)
	act.Triggered().Attach(trigFunc)
	if icoPath != "" {
		ico, _ := walk.NewIconFromFile(icoPath)
		act.SetImage(ico)
	}
	wm.tb.Actions().Add(act)
	return act
}

/**
*	AddToolSeparator
**/
func (wm *WalkUI) AddToolSeparator() *walk.Action {
	act := walk.NewSeparatorAction()
	wm.tb.Actions().Add(act)
	return act
}

/**
*	ToolBarEnd
**/
func (wm *WalkUI) ToolBarEnd() {
	wm.Window().SetToolBar(wm.tb)
}

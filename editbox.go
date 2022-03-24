package walkmgr

import (
	"github.com/pirogom/walk"
	"github.com/pirogom/win"
)

/**
*	NumberEdit
**/
func (wm *walk_ui) NumberEdit(val ...int) *walk.NumberEdit {
	ne, _ := walk.NewNumberEdit(wm.Parent())
	if len(val) > 0 {
		ne.SetValue(float64(val[0]))
	}
	wm.Append(ne)
	return ne
}

/**
*	LineEdit
**/
func (wm *walk_ui) LineEdit(msg ...string) *walk.LineEdit {
	ne, _ := walk.NewLineEdit(wm.Parent())
	if len(msg) > 0 {
		ne.SetText(msg[0])
	}
	wm.Append(ne)
	return ne
}

/**
*	LineStatic
**/
func (wm *walk_ui) LineStatic(msg ...string) *walk.LineEdit {
	ne, _ := walk.NewLineEdit(wm.Parent())
	if len(msg) > 0 {
		ne.SetText(msg[0])
	}
	ne.SetReadOnly(true)
	wm.Append(ne)
	return ne
}

/**
*	TextEdit
**/
func (wm *walk_ui) TextEdit(msg ...string) *walk.TextEdit {
	ne, _ := walk.NewTextEdit(wm.Parent())
	if len(msg) > 0 {
		ne.SetText(msg[0])
	}
	wm.Append(ne)
	return ne
}

/**
*	TextStatic
**/
func (wm *walk_ui) TextStatic(msg ...string) *walk.TextEdit {
	ne, _ := walk.NewTextEdit(wm.Parent())
	if len(msg) > 0 {
		ne.SetText(msg[0])
	}
	ne.SetReadOnly(true)
	wm.Append(ne)
	return ne
}

/**
*	TextArea
**/
func (wm *walk_ui) TextArea(msg ...string) *walk.TextEdit {
	ne, _ := walk.NewTextEditWithStyle(wm.Parent(), win.WS_VSCROLL)
	if len(msg) > 0 {
		ne.SetText(msg[0])
	}
	wm.Append(ne)
	return ne
}

/**
*	TextAreaStatic
**/
func (wm *walk_ui) TextAreaStatic(msg ...string) *walk.TextEdit {
	ne, _ := walk.NewTextEditWithStyle(wm.Parent(), win.WS_VSCROLL)
	if len(msg) > 0 {
		ne.SetText(msg[0])
	}
	ne.SetReadOnly(true)
	wm.Append(ne)
	return ne
}

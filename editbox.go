package walkmgr

import (
	"github.com/pirogom/walk"
	"github.com/pirogom/win"
)

/**
*	NumberEdit
**/
func (wm *walkmgr) NumberEdit(val ...int) *walk.NumberEdit {
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
func (wm *walkmgr) LineEdit(msg ...string) *walk.LineEdit {
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
func (wm *walkmgr) LineStatic(msg ...string) *walk.LineEdit {
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
func (wm *walkmgr) TextEdit(msg ...string) *walk.TextEdit {
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
func (wm *walkmgr) TextStatic(msg ...string) *walk.TextEdit {
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
func (wm *walkmgr) TextArea(msg ...string) *walk.TextEdit {
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
func (wm *walkmgr) TextAreaStatic(msg ...string) *walk.TextEdit {
	ne, _ := walk.NewTextEditWithStyle(wm.Parent(), win.WS_VSCROLL)
	if len(msg) > 0 {
		ne.SetText(msg[0])
	}
	ne.SetReadOnly(true)
	wm.Append(ne)
	return ne
}

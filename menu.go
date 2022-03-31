package walkmgr

import "github.com/pirogom/walk"

type MenuMgr struct {
	Menu    *walk.Menu
	MenuAct *walk.Action
}

/**
*	NewMenu
**/
func NewMenu(text string) *MenuMgr {
	m := MenuMgr{}
	m.Menu, _ = walk.NewMenu()
	m.MenuAct = walk.NewMenuAction(m.Menu)
	m.MenuAct.SetText(text)
	return &m
}

/**
*	AddAction
**/
func (m *MenuMgr) AddAction(text string, trigerFunc func()) *walk.Action {
	act := walk.NewAction()
	act.SetText(text)
	act.Triggered().Attach(trigerFunc)
	m.Menu.Actions().Add(act)
	return act
}

/**
*	AddCheckAction
**/
func (m *MenuMgr) AddCheck(text string, checked bool, trigerFunc func()) *walk.Action {
	act := walk.NewAction()
	act.SetText(text)

	act.SetCheckable(true)
	act.SetChecked(checked)
	act.Triggered().Attach(trigerFunc)

	m.Menu.Actions().Add(act)
	return act
}

/**
*	AddSeparator
**/
func (m *MenuMgr) AddSeparator() *walk.Action {
	act := walk.NewSeparatorAction()
	m.Menu.Actions().Add(act)
	return act
}

/**
*	AddMenu
**/
func (m *MenuMgr) AddMenu(in *MenuMgr) *MenuMgr {
	if m == in {
		panic("must m != in")
	}
	m.Menu.Actions().Add(in.MenuAct)
	return m
}

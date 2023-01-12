package walkmgr

import (
	"github.com/pirogom/walk"
)

/**
*	TreeView
**/
type TreeView struct {
	wm *WalkUI
	tv *walk.TreeView
	tm *TreeModel
}

/**
*	TreeViewItem
**/
type TreeViewItem struct {
	name     string
	parent   *TreeViewItem
	children []*TreeViewItem
	icon     *walk.Icon
}

/**
*	AddItem
**/
func (d *TreeViewItem) AddItem(name string, icon *walk.Icon) *TreeViewItem {
	nd := new(TreeViewItem)
	nd.name = name
	if icon != nil {
		nd.icon = icon
	}
	nd.parent = d
	d.children = append(d.children, nd)
	return nd
}

/**
*	Text
**/
func (d *TreeViewItem) Text() string {
	return d.name
}

/**
*	Parent
**/
func (d *TreeViewItem) Parent() walk.TreeItem {
	if d.parent == nil {
		return nil
	}

	return d.parent
}

/**
*	ChildCount
**/
func (d *TreeViewItem) ChildCount() int {
	if d.children == nil {
		return 0
	}
	return len(d.children)
}

/**
*	ChildAt
**/
func (d *TreeViewItem) ChildAt(index int) walk.TreeItem {
	return d.children[index]
}

/**
*	Image
**/
func (d *TreeViewItem) Image() interface{} {
	if d.icon != nil {
		return d.icon
	}
	return nil
}

/**
* ResetChildren
**/
func (d *TreeViewItem) ResetChildren() error {
	d.children = nil
	return nil
}

/**
*	TreeModel
**/
type TreeModel struct {
	walk.TreeModelBase
	roots []*TreeViewItem
}

/**
*	LazyPopulation
**/
func (*TreeModel) LazyPopulation() bool {
	return true
}

/**
*	RootCount
**/
func (m *TreeModel) RootCount() int {
	return len(m.roots)
}

/**
*	RootAt
**/
func (m *TreeModel) RootAt(index int) walk.TreeItem {
	return m.roots[index]
}

/**
*	NewTreeView
**/
func (wm *WalkUI) NewTreeView() *TreeView {
	nd := TreeView{}

	nd.tv, _ = walk.NewTreeView(wm.Parent())
	nd.wm = wm
	nd.tm = new(TreeModel)
	return &nd
}

/**
*	AddItem
**/
func (t *TreeView) AddItem(name string, icon *walk.Icon) *TreeViewItem {
	nd := new(TreeViewItem)
	nd.name = name
	if icon != nil {
		nd.icon = icon
	}
	t.tm.roots = append(t.tm.roots, nd)
	return nd
}

/**
*	Create
**/
func (t *TreeView) Create() {
	t.tv.SetModel(t.tm)
	t.wm.Append(t.tv)
}

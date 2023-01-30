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
	data     interface{}
}

/**
*	Data
**/
func (d *TreeViewItem) Data() interface{} {
	return d.data
}

/**
*	SetData
**/
func (d *TreeViewItem) SetData(data interface{}) {
	d.data = data
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

	nd.tv.SetModel(nd.tm)
	nd.wm.Append(nd.tv)

	return &nd
}

/**
*	Clear
**/
func (t *TreeView) Clear() {
	t.tm = nil
	t.tm = new(TreeModel)
	t.tv.SetModel(t.tm)
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
*	GetTreeView
**/
func (t *TreeView) GetTreeView() *walk.TreeView {
	return t.tv
}

/**
*	UpdateItems
**/
func (t *TreeView) UpdateItems() error {
	return t.tv.UpdateItems()
}

/**
*	CurrentItem
**/
func (t *TreeView) CurrentItem() *TreeViewItem {
	return t.tv.CurrentItem().(*TreeViewItem)
}

/**
*	SetCurrentItem
**/
func (t *TreeView) SetCurrentItem(ti *TreeViewItem) error {
	return t.tv.SetCurrentItem(ti)
}

/**
*	CurrentItemChanged
**/
func (t *TreeView) CurrentItemChanged(f func()) int {
	return t.tv.CurrentItemChanged().Attach(f)
}

/**
*	SetExpanded
**/
func (t *TreeView) SetExpanded(item *TreeViewItem, onoff bool) error {
	return t.tv.SetExpanded(item, onoff)
}

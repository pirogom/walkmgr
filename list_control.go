package walkmgr

import (
	"sort"

	"github.com/pirogom/walk"
)

/**
*	ListControlItem
**/
type ListControlItem struct {
	values  []string
	checked bool
}

/**
*	ListControlModel
**/
type ListControlModel struct {
	walk.TableModelBase
	items []ListControlItem
}

/**
*	RowCount
**/
func (m *ListControlModel) RowCount() int {
	return len(m.items)
}

/**
*	Value
**/
func (m *ListControlModel) Value(row, col int) interface{} {
	if row < 0 || row >= len(m.items) {
		panic("bad row idx")
	}
	if col < 0 || col >= len(m.items[row].values) {
		panic("bad col idx")
	}

	return m.items[row].values[col]
}

/**
*	Checked
**/
func (m *ListControlModel) Checked(row int) bool {
	return m.items[row].checked
}

/**
*	CheckedCount
**/
func (m *ListControlModel) CheckedCount() int {
	var cnt int
	for _, item := range m.items {
		if item.checked {
			cnt++
		}
	}
	return cnt
}

/**
*	SetChecked
**/
func (m *ListControlModel) SetChecked(row int, checked bool) error {
	m.items[row].checked = checked
	return nil
}

/**
*	ResetRows
**/
func (m *ListControlModel) ResetRows() {
	m.items = nil
	m.PublishRowsReset()
}

type ListControlCfg struct {
	CheckBox    bool
	MultiSelect bool
}

type ListCtrlAlign int32
type ListCtrlOrderState bool

const (
	LISTCTRL_ALIGN_RIGHT  ListCtrlAlign = 0
	LISTCTRL_ALIGN_LEFT   ListCtrlAlign = 1
	LISTCTRL_ALIGN_CENTER ListCtrlAlign = 2

	LS_ORDER_ASC  ListCtrlOrderState = false
	LS_ORDER_DESC ListCtrlOrderState = true
)

type ListControlColumn struct {
	Title      string
	Width      int
	Align      ListCtrlAlign
	Order      bool
	orderState ListCtrlOrderState
}

/**
*	ListControl
**/
type ListControl struct {
	cbModel *ListControlModel
	tv      *walk.TableView
	wm      *WalkUI
	th      []ListControlColumn
	cfg     ListControlCfg
}

/**
*	NewListControl
**/
func NewListControl(wm *WalkUI, cfg *ListControlCfg) *ListControl {
	nd := ListControl{}
	nd.cbModel = new(ListControlModel)
	nd.wm = wm
	if cfg != nil {
		nd.cfg = *cfg
	}
	return &nd
}

/**
*	AddColumn
**/
func (t *ListControl) AddColumn(title string, width int, align ListCtrlAlign, order bool) {
	od := ListControlColumn{Title: title, Width: width, Align: align, Order: order}
	t.th = append(t.th, od)
}

/**
*	Create
**/
func (t *ListControl) Create() {
	t.tv, _ = walk.NewTableView(t.wm.Parent())
	t.tv.SetCheckBoxes(t.cfg.CheckBox)
	t.tv.SetMultiSelection(t.cfg.MultiSelect)
	t.tv.SetModel(t.cbModel)

	for i := 0; i < len(t.th); i++ {
		col := walk.NewTableViewColumn()
		col.SetTitle(t.th[i].Title)
		col.SetWidth(t.th[i].Width)

		switch t.th[i].Align {
		case LISTCTRL_ALIGN_CENTER:
			col.SetAlignment(walk.AlignCenter)
		case LISTCTRL_ALIGN_RIGHT:
			col.SetAlignment(walk.AlignFar)
		case LISTCTRL_ALIGN_LEFT:
			col.SetAlignment(walk.AlignNear)
		}
		t.tv.Columns().Add(col)
	}
	t.wm.Append(t.tv)

	t.registEvent()
}

/**
*	registEvent
**/
func (t *ListControl) registEvent() {
	t.tv.ColumnClicked().Attach(t.columnOrderingEvent)
}

/**
*	columnOrderingEvent
**/
func (t *ListControl) columnOrderingEvent(col int) {
	if t.th[col].orderState == LS_ORDER_ASC {
		t.th[col].orderState = LS_ORDER_DESC
	} else {
		t.th[col].orderState = LS_ORDER_ASC
	}

	keys := []string{}
	sortMap := make(map[string]ListControlItem)

	for _, item := range t.cbModel.items {
		keys = append(keys, item.values[col])
		sortMap[item.values[col]] = item
	}

	if t.th[col].orderState == LS_ORDER_ASC {
		sort.Strings(keys)
	} else {
		sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	}

	t.cbModel.ResetRows()

	for _, k := range keys {
		t.cbModel.items = append(t.cbModel.items, sortMap[k])
	}
	t.cbModel.PublishRowsReset()
}

/**
*	GetItemCount
**/
func (t *ListControl) GetItemCount() int {
	return t.cbModel.RowCount()
}

/**
*	GetItemData
**/
func (t *ListControl) GetItemData(row, col int) string {
	if row < 0 || row >= len(t.cbModel.items) {
		panic("bad row idx")
	}
	if col < 0 || col >= len(t.cbModel.items[row].values) {
		panic("bad col idx")
	}
	return t.cbModel.items[row].values[col]
}

/**
*	SetItemData
**/
func (t *ListControl) SetItemData(row, col int, value string) {
	if row < 0 || row >= len(t.cbModel.items) {
		panic("bad row idx")
	}
	if col < 0 || col >= len(t.cbModel.items[row].values) {
		panic("bad col idx")
	}
	t.cbModel.items[row].values[col] = value
}

/**
*	RemoveAll
**/
func (t *ListControl) RemoveAll() {
	t.cbModel.ResetRows()
}

/**
*	UpdateItem
**/
func (t *ListControl) UpdateItem(row int) error {
	if row < 0 || row >= len(t.cbModel.items) {
		panic("bad row idx")
	}
	return t.tv.UpdateItem(row)
}

/**
*	UpdateAll
**/
func (t *ListControl) UpdateAll() {
	for i := range t.cbModel.items {
		t.UpdateItem(i)
	}
}

/**
*	DeleteItem
**/
func (t *ListControl) DeleteItem(row int) {
	if row < 0 || row >= len(t.cbModel.items) {
		panic("bad row idx")
	}
	t.cbModel.items = append(t.cbModel.items[:row], t.cbModel.items[row+1:]...)
	t.cbModel.PublishRowsRemoved(row, row+1)
}

/**
*	AddItemData
**/
func (t *ListControl) AddItemData(value string) int {
	ni := ListControlItem{}
	ni.values = make([]string, len(t.th))
	ni.values[0] = value
	rowCnt := t.cbModel.RowCount()
	t.cbModel.items = append(t.cbModel.items, ni)
	t.cbModel.PublishRowsInserted(rowCnt, rowCnt)
	return rowCnt
}

/**
*	ListControl
**/
func (t *ListControl) SelectedItem() *ListControlItem {
	idx := t.tv.CurrentIndex()
	if idx == -1 {
		return nil
	}
	return &t.cbModel.items[idx]
}

/**
*	SelectedItemIndex
**/
func (t *ListControl) SelectedItemIndex() int {
	return t.tv.CurrentIndex()
}

/**
*	AllSelectedItem
**/
func (t *ListControl) AllSelectedItem() []ListControlItem {
	li := []ListControlItem{}
	idxs := t.tv.SelectedIndexes()
	for _, idx := range idxs {
		li = append(li, t.cbModel.items[idx])
	}
	return li
}

/**
*	AllSelectedItemIndex
**/
func (t *ListControl) AllSelectedItemIndex() []int {
	return t.tv.SelectedIndexes()
}

/**
*	SelectedItemCount
**/
func (t *ListControl) SelectedItemCount() int {
	return len(t.tv.SelectedIndexes())
}

/**
*	CheckedAll
**/
func (t *ListControl) CheckedAll(checked bool) {
	if !t.tv.CheckBoxes() {
		return
	}
	for itemIdx, _ := range t.cbModel.items {
		t.Checked(itemIdx, checked)
	}
}

/**
*	Checked
**/
func (t *ListControl) Checked(idx int, checked bool) {
	t.cbModel.SetChecked(idx, checked)
	t.cbModel.PublishRowChanged(idx)
}

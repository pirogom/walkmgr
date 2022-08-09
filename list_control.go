package walkmgr

import "github.com/pirogom/walk"

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

type lcCheckOpt bool
type lcSelectOpt bool

const (
	LISTCONTROL_NO_CHECKBOX     lcCheckOpt  = false
	LISTCONTROL_CHECKBOX        lcCheckOpt  = true
	LISTCONTROL_NO_MULTI_SELECT lcSelectOpt = false
	LISTCONTROL_MULTI_SELECT    lcSelectOpt = true
)

/**
*	ListControl
**/
type ListControl struct {
	cbModel        *ListControlModel
	tv             *walk.TableView
	wm             *WalkUI
	th             TableHeader
	hasCheckbox    lcCheckOpt
	useMultiSelect lcSelectOpt
}

/**
*	NewListControl
**/
func NewListControl(wm *WalkUI, checkbox lcCheckOpt, multiselect lcSelectOpt) *ListControl {
	nd := ListControl{}
	nd.cbModel = new(ListControlModel)
	nd.wm = wm
	nd.hasCheckbox = checkbox
	nd.useMultiSelect = multiselect
	return &nd
}

/**
*	AddColumn
**/
func (t *ListControl) AddColumn(title string, width int, align ...string) {
	var av string
	if len(align) > 0 {
		av = align[0]
	} else {
		av = "left"
	}
	t.th.Header = append(t.th.Header, TvHeader{Title: title, Width: width, Align: av})
}

/**
*	Create
**/
func (t *ListControl) Create() {
	t.tv = t.wm.TableView(t.cbModel, t.th.Get(), bool(t.hasCheckbox), bool(t.useMultiSelect))
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
	ni.values = make([]string, len(t.th.Header))
	ni.values[0] = value
	rowCnt := t.cbModel.RowCount()
	t.cbModel.items = append(t.cbModel.items, ni)
	t.cbModel.PublishRowsInserted(rowCnt, rowCnt)
	return rowCnt
}

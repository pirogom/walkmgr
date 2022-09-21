package walkmgr

import (
	"sort"

	"github.com/pirogom/walk"
)

type ItemDoubleClickedFunc func(idx int)
type CurrentIndexChangedFunc func(idx int)
type KeyDownFunc func(key walk.Key)

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

	itemDblClickedProc ItemDoubleClickedFunc
	itemChangedProc    CurrentIndexChangedFunc
	keyDownProc        KeyDownFunc
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
	t.tv.ItemActivated().Attach(t.itemDblClickedEvent)
	t.tv.MouseWheel().Attach(t.mouseWheelEvent)
	t.tv.CurrentIndexChanged().Attach(t.currItemIndexChangedEvent)
	t.tv.KeyDown().Attach(t.keydownEvnet)
}

/**
*	keydownEvnet
**/
func (t *ListControl) keydownEvnet(key walk.Key) {
	defer func() {
		if t.keyDownProc != nil {
			t.keyDownProc(key)
		}
	}()

	if key == walk.KeySpace {
		currIdx := t.tv.CurrentIndex()
		if currIdx > -1 {
			selectItem := t.tv.SelectedIndexes()

			if len(selectItem) == 0 {
				return
			}

			checked := t.cbModel.items[currIdx].checked

			for _, itemVal := range selectItem {

				if itemVal != currIdx {
					t.cbModel.items[itemVal].checked = checked
					t.tv.UpdateItem(itemVal)
				}
			}
		}
	}
}

/**
*	currItemIndexChangedEvent
**/
func (t *ListControl) currItemIndexChangedEvent() {
	nCurrIten := t.tv.CurrentIndex()

	if nCurrIten > -1 && t.itemChangedProc != nil {
		t.itemChangedProc(nCurrIten)
	}
}

/**
*	mouseWheelEvent
**/
func (t *ListControl) mouseWheelEvent(x, y int, button walk.MouseButton) {
	wDelta := walk.MouseWheelEventDelta(button)

	if wDelta > 0 {
		t.CurrentIndexUp()
	} else {
		t.CurrentIndexDown()
	}
}

/**
*	itemDblClickedEvent
**/
func (t *ListControl) itemDblClickedEvent() {
	nCurrItem := t.tv.CurrentIndex()

	if t.tv.CheckBoxes() {
		if nCurrItem > -1 {
			if t.cbModel.items[nCurrItem].checked {
				t.cbModel.items[nCurrItem].checked = false
			} else {
				t.cbModel.items[nCurrItem].checked = true
			}
			t.cbModel.PublishRowChanged(nCurrItem)
		}
	}
	if nCurrItem > -1 && t.itemDblClickedProc != nil {
		t.itemDblClickedProc(nCurrItem)
	}
}

/**
*	columnOrderingEvent
**/
func (t *ListControl) columnOrderingEvent(col int) {
	if !t.th[col].Order {
		return
	}

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

/**
*	CheckedCount
**/
func (t *ListControl) CheckedCount() int {
	if !t.tv.CheckBoxes() {
		return 0
	}
	return t.cbModel.CheckedCount()
}

/**
*	CheckedIndexs
**/
func (t *ListControl) CheckedIndexs() []int {
	ci := []int{}

	for i := range t.cbModel.items {
		if t.cbModel.items[i].checked {
			ci = append(ci, i)
		}
	}

	return ci
}

/**
*	ItemDoubleClicked
**/
func (t *ListControl) ItemDoubleClicked(afterFunc ItemDoubleClickedFunc) {
	t.itemDblClickedProc = afterFunc
}

/**
*	CurrIndexChanged
**/
func (t *ListControl) CurrIndexChanged(afterFunc CurrentIndexChangedFunc) {
	t.itemChangedProc = afterFunc
}

/**
*	KeyDown
**/
func (t *ListControl) KeyDown(afterFunc KeyDownFunc) {
	t.keyDownProc = afterFunc
}

/**
*	CurrentIndexUp
**/
func (t *ListControl) CurrentIndexUp() {
	if t.cbModel.RowCount() == 0 {
		return
	}
	idx := t.tv.CurrentIndex()

	if idx == 0 {
		return
	} else if idx == -1 {
		t.tv.SetCurrentIndex(0)
	} else {
		t.tv.SetCurrentIndex(idx - 1)
	}
}

/**
*	CurrentIndexDown
**/
func (t *ListControl) CurrentIndexDown() {
	if t.cbModel.RowCount() == 0 {
		return
	}
	idx := t.tv.CurrentIndex()
	if idx+1 >= t.cbModel.RowCount() {
		return
	} else {
		t.tv.SetCurrentIndex(idx + 1)
	}
}

/**
*	CurrentItemChecked
**/
func (t *ListControl) CurrentItemCheck() (int, bool) {
	if !t.tv.CheckBoxes() {
		return -1, false
	}
	idx := t.tv.CurrentIndex()
	if idx > -1 {
		if t.cbModel.items[idx].checked {
			t.cbModel.items[idx].checked = false
		} else {
			t.cbModel.items[idx].checked = true
		}
		t.cbModel.PublishRowChanged(idx)
		return idx, t.cbModel.items[idx].checked
	}
	return -1, false
}

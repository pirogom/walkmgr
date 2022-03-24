package walkmgr

import "github.com/pirogom/walk"

/**
*	TvHeader
**/
type TvHeader struct {
	Title string
	Width int
	Align string
}

type TableHeader struct {
	Header []TvHeader
}

func NewTH(title string, width int) *TableHeader {
	th := TableHeader{}
	th.Add(title, width)
	return &th
}

func (t *TableHeader) Add(title string, width int) *TableHeader {
	t.Header = append(t.Header, TvHeader{Title: title, Width: width})
	return t
}

func (t *TableHeader) Get() []TvHeader {
	return t.Header
}

/**
*	TableView
**/
func (wm *WalkUI) TableView(model interface{}, header []TvHeader, checkBox bool, multiSelect bool) *walk.TableView {
	tv, _ := walk.NewTableView(wm.Parent())
	tv.SetCheckBoxes(checkBox)
	tv.SetMultiSelection(multiSelect)
	tv.SetModel(model)

	for i := 0; i < len(header); i++ {
		col := walk.NewTableViewColumn()
		col.SetTitle(header[i].Title)
		col.SetWidth(header[i].Width)

		switch header[i].Align {
		case "center":
			col.SetAlignment(walk.AlignCenter)
		case "right":
			col.SetAlignment(walk.AlignFar)
		case "left":
			col.SetAlignment(walk.AlignNear)
		}
		tv.Columns().Add(col)
	}
	wm.Append(tv)
	return tv
}

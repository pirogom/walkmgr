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

// type TableHeader struct {
// 	Header []TvHeader
// }

// func (t *TableHeader) Append(title string, width int, height int) *TableHeader {

// 	t.Header = Append(t.Header, TvHeader{Title: title, Width:width})

// }

/**
*	TableView
**/
func (wm *walkmgr) TableView(model interface{}, header []TvHeader, checkBox bool, multiSelect bool) *walk.TableView {

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

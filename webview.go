package walkmgr

import "github.com/pirogom/walk"

/**
*	WebView
**/
func (wm *walkmgr) WebView(url string) *walk.WebView {
	wv, _ := walk.NewWebView(wm.Parent())
	wv.SetURL(url)

	wm.Append(wv)
	return wv
}

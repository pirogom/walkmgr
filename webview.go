package walkmgr

import "github.com/pirogom/walk"

/**
*	WebView
**/
func (wm *WalkUI) WebView(url string) *walk.WebView {
	wv, _ := walk.NewWebView(wm.Parent())
	wv.SetURL(url)
	wv.SetSilent()
	wm.Append(wv)
	return wv
}

/**
*	WebViewWithAlert
**/
func (wm *WalkUI) WebViewWithAlert(url string) *walk.WebView {
	wv, _ := walk.NewWebView(wm.Parent())
	wv.SetURL(url)
	wm.Append(wv)
	return wv
}

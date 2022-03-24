package walkmgr

import "github.com/pirogom/walk"

/**
*	WebView
**/
func (wm *walk_ui) WebView(url string) *walk.WebView {
	wv, _ := walk.NewWebView(wm.Parent())
	wv.SetURL(url)
	wv.SetSilent()
	wm.Append(wv)
	return wv
}

/**
*	WebViewWithAlert
**/
func (wm *walk_ui) WebViewWithAlert(url string) *walk.WebView {
	wv, _ := walk.NewWebView(wm.Parent())
	wv.SetURL(url)
	wm.Append(wv)
	return wv
}

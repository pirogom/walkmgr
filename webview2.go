package walkmgr

import (
	"errors"

	"github.com/jchv/go-webview2/pkg/edge"
	"github.com/pirogom/walk"
)

var Webview2InitErr = errors.New("Webview2 object not inited")

/**
 * 	WebView2
 */
type WebView2 struct {
	iv     *walk.ImageView
	wv     *edge.Chromium
	inited bool
}

/**
 * IsInited
 */
func (wv2 *WebView2) IsInited() bool {
	return wv2.inited
}

/**
 * 	SizeChanged
 */
func (wv2 *WebView2) SizeChanged() {
	if wv2.inited {
		wv2.wv.NotifyParentWindowPositionChanged()
		wv2.wv.Resize()
	}
}

/**
 * 	Navigate
 */
func (wv2 *WebView2) Navigate(url string) {
	if wv2.inited {
		wv2.wv.Navigate(url)
	}
}

/**
 * 	NavigateToString
 */
func (wv2 *WebView2) NavigateToString(html string) {
	if wv2.inited {
		wv2.wv.NavigateToString(html)
	}
}

/**
 * get_IsScriptEnabled
 *
 * https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings?view=webview2-1.0.1901.177#get_isscriptenabled
 */
func (wv2 *WebView2) GetIsScriptEnabled() (bool, error) {
	if !wv2.inited {
		return false, Webview2InitErr
	}

	cfg, err := wv2.wv.GetSettings()

	if err != nil {
		return false, err
	}

	return cfg.GetIsScriptEnabled()
}

/**
 * put_IsScriptEnabled
 *
 * https://learn.microsoft.com/en-us/microsoft-edge/webview2/reference/win32/icorewebview2settings?view=webview2-1.0.1901.177#put_isscriptenabled
 */
func (wv2 *WebView2) PutIsScriptEnabled(isScriptEnabled bool) error {
	if !wv2.inited {
		return Webview2InitErr
	}

	cfg, err := wv2.wv.GetSettings()

	if err != nil {
		return err
	}
	return cfg.PutIsScriptEnabled(isScriptEnabled)
}

/**
 * get_IsWebMessageEnabled
 */
func (wv2 *WebView2) GetIsWebMessageEnabled() (bool, error) {
	if !wv2.inited {
		return false, Webview2InitErr
	}

	cfg, err := wv2.wv.GetSettings()

	if err != nil {
		return false, err
	}

	return cfg.GetIsWebMessageEnabled()
}

/**
 * put_IsWebMessageEnabled
 */
func (wv2 *WebView2) PutIsWebMessageEnabled(isWebMessageEnabled bool) error {
	if !wv2.inited {
		return Webview2InitErr
	}

	cfg, err := wv2.wv.GetSettings()

	if err != nil {
		return err
	}

	return cfg.PutIsWebMessageEnabled(isWebMessageEnabled)
}

/**
 * get_AreDefaultScriptDialogsEnabled
 */
func (wv2 *WebView2) GetAreDefaultScriptDialogsEnabled() (bool, error) {
	if !wv2.inited {
		return false, Webview2InitErr
	}

	cfg, err := wv2.wv.GetSettings()

	if err != nil {
		return false, err
	}
	return cfg.GetAreDefaultScriptDialogsEnabled()
}

/**
 * put_AreDefaultScriptDialogsEnabled
 */
func (wv2 *WebView2) PutAreDefaultScriptDialogsEnabled(areDefaultScriptDialogsEnabled bool) error {
	if !wv2.inited {
		return Webview2InitErr
	}

	cfg, err := wv2.wv.GetSettings()

	if err != nil {
		return err
	}
	return cfg.PutAreDefaultScriptDialogsEnabled(areDefaultScriptDialogsEnabled)
}

/**
 * get_IsStatusBarEnabled
 */
func (wv2 *WebView2) GetIsStatusBarEnabled() (bool, error) {
	if !wv2.inited {
		return false, Webview2InitErr
	}

	cfg, err := wv2.wv.GetSettings()

	if err != nil {
		return false, err
	}

	return cfg.GetIsStatusBarEnabled()
}

/**
 * put_IsStatusBarEnabled
 */
func (wv2 *WebView2) PutIsStatusBarEnabled(isStatusBarEnabled bool) error {
	if !wv2.inited {
		return Webview2InitErr
	}

	cfg, err := wv2.wv.GetSettings()

	if err != nil {
		return err
	}
	return cfg.PutIsStatusBarEnabled(isStatusBarEnabled)
}

/**
 * get_AreDevToolsEnabled
 */
func (wv2 *WebView2) GetAreDevToolsEnabled() (bool, error) {
	if !wv2.inited {
		return false, Webview2InitErr
	}

	cfg, err := wv2.wv.GetSettings()

	if err != nil {
		return false, err
	}

	return cfg.GetAreDevToolsEnabled()
}

/**
 * put_AreDevToolsEnabled
 */
func (wv2 *WebView2) PutAreDevToolsEnabled(areDevToolsEnabled bool) error {
	if !wv2.inited {
		return Webview2InitErr
	}

	cfg, err := wv2.wv.GetSettings()

	if err != nil {
		return err
	}
	return cfg.PutAreDevToolsEnabled(areDevToolsEnabled)
}

/**
 * get_AreDefaultContextMenusEnabled
 */
func (wv2 *WebView2) GetAreDefaultContextMenusEnabled() (bool, error) {
	if !wv2.inited {
		return false, Webview2InitErr
	}

	cfg, err := wv2.wv.GetSettings()

	if err != nil {
		return false, err
	}
	return cfg.GetAreDefaultContextMenusEnabled()
}

/**
 * put_AreDefaultContextMenusEnabled
 */
func (wv2 *WebView2) PutAreDefaultContextMenusEnabled(enabled bool) error {
	if !wv2.inited {
		return Webview2InitErr
	}

	cfg, err := wv2.wv.GetSettings()

	if err != nil {
		return err
	}
	return cfg.PutAreDefaultContextMenusEnabled(enabled)
}

/**
 * get_AreHostObjectsAllowed
 */
func (wv2 *WebView2) GetAreHostObjectsAllowed() (bool, error) {
	if !wv2.inited {
		return false, Webview2InitErr
	}

	cfg, err := wv2.wv.GetSettings()

	if err != nil {
		return false, err
	}
	return cfg.GetAreHostObjectsAllowed()
}

/**
 * put_AreHostObjectsAllowed
 */
func (wv2 *WebView2) PutAreHostObjectsAllowed(allowed bool) error {
	if !wv2.inited {
		return Webview2InitErr
	}

	cfg, err := wv2.wv.GetSettings()

	if err != nil {
		return err
	}
	return cfg.PutAreHostObjectsAllowed(allowed)
}

/**
 * get_IsZoomControlEnabled
 */
func (wv2 *WebView2) GetIsZoomControlEnabled() (bool, error) {
	if !wv2.inited {
		return false, Webview2InitErr
	}

	cfg, err := wv2.wv.GetSettings()

	if err != nil {
		return false, err
	}
	return cfg.GetIsZoomControlEnabled()
}

/**
 * put_IsZoomControlEnabled
 */
func (wv2 *WebView2) PutIsZoomControlEnabled(enabled bool) error {
	if !wv2.inited {
		return Webview2InitErr
	}

	cfg, err := wv2.wv.GetSettings()

	if err != nil {
		return err
	}
	return cfg.PutIsZoomControlEnabled(enabled)
}

/**
 * get_IsBuiltInErrorPageEnabled
 */
func (wv2 *WebView2) GetIsBuiltInErrorPageEnabled() (bool, error) {
	if !wv2.inited {
		return false, Webview2InitErr
	}

	cfg, err := wv2.wv.GetSettings()

	if err != nil {
		return false, err
	}
	return cfg.GetIsBuiltInErrorPageEnabled()
}

/**
 * put_IsBuiltInErrorPageEnabled
 */
func (wv2 *WebView2) PutIsBuiltInErrorPageEnabled(enabled bool) error {
	if !wv2.inited {
		return Webview2InitErr
	}

	cfg, err := wv2.wv.GetSettings()

	if err != nil {
		return err
	}
	return cfg.PutIsBuiltInErrorPageEnabled(enabled)
}

/**
 * get User-Agent string
 */
func (wv2 *WebView2) GetUserAgent() (string, error) {
	if !wv2.inited {
		return "", Webview2InitErr
	}

	cfg, err := wv2.wv.GetSettings()

	if err != nil {
		return "", err
	}

	return cfg.GetUserAgent()
}

/**
 * Set User-Agent string
 */
func (wv2 *WebView2) PutUserAgent(userAgent string) error {
	if !wv2.inited {
		return Webview2InitErr
	}

	cfg, err := wv2.wv.GetSettings()

	if err != nil {
		return err
	}
	return cfg.PutUserAgent(userAgent)
}

/**
 * get_AreBrowserAcceleratorKeysEnabled
 */
func (wv2 *WebView2) GetAreBrowserAcceleratorKeysEnabled() (bool, error) {
	if !wv2.inited {
		return false, Webview2InitErr
	}

	cfg, err := wv2.wv.GetSettings()

	if err != nil {
		return false, err
	}
	return cfg.GetAreBrowserAcceleratorKeysEnabled()
}

/**
 * put_AreBrowserAcceleratorKeysEnabled
 */
func (wv2 *WebView2) PutAreBrowserAcceleratorKeysEnabled(enabled bool) error {
	if !wv2.inited {
		return Webview2InitErr
	}

	cfg, err := wv2.wv.GetSettings()

	if err != nil {
		return err
	}
	return cfg.PutAreBrowserAcceleratorKeysEnabled(enabled)
}

/**
 * get_IsPinchZoomEnabled
 */
func (wv2 *WebView2) GetIsPinchZoomEnabled() (bool, error) {
	if !wv2.inited {
		return false, Webview2InitErr
	}

	cfg, err := wv2.wv.GetSettings()

	if err != nil {
		return false, err
	}
	return cfg.GetIsPinchZoomEnabled()
}

/**
 * put_IsPinchZoomEnabled
 */
func (wv2 *WebView2) PutIsPinchZoomEnabled(enabled bool) error {
	if !wv2.inited {
		return Webview2InitErr
	}

	cfg, err := wv2.wv.GetSettings()

	if err != nil {
		return err
	}
	return cfg.PutIsPinchZoomEnabled(enabled)
}

/**
 * get_IsSwipeNavigationEnabled
 */
func (wv2 *WebView2) GetIsSwipeNavigationEnabled() (bool, error) {
	if !wv2.inited {
		return false, Webview2InitErr
	}

	cfg, err := wv2.wv.GetSettings()

	if err != nil {
		return false, err
	}
	return cfg.GetIsSwipeNavigationEnabled()
}

/**
 * put_IsSwipeNavigationEnabled
 */
func (wv2 *WebView2) PutIsSwipeNavigationEnabled(enabled bool) error {
	if !wv2.inited {
		return Webview2InitErr
	}

	cfg, err := wv2.wv.GetSettings()

	if err != nil {
		return err
	}
	return cfg.PutIsSwipeNavigationEnabled(enabled)
}

/**
*	WebView2
**/
func (wm *WalkUI) WebView2(dataPath string, succssFunc func(), failFunc func()) *WebView2 {
	wv2 := WebView2{}

	wv2.wv = edge.NewChromium()
	wv2.wv.SetPermission(edge.CoreWebView2PermissionKindClipboardRead, edge.CoreWebView2PermissionStateAllow)
	wv2.wv.DataPath = dataPath

	wv2.iv = wm.ImageView(walk.ImageViewModeStretch)

	wv2.iv.SizeChanged().Attach(func() {
		wv2.SizeChanged()
	})

	id := webview2InitData{}
	id.wv2 = &wv2
	id.successFunc = succssFunc
	id.failFunc = failFunc

	wm.webview2InitFunctions = append(wm.webview2InitFunctions, id)
	
	return &wv2
}

/*
*
webview2InitProc
*/
func (wm *WalkUI) webview2InitProc() {
	for i := 0; i < len(wm.webview2InitFunctions); i++ {
		wm.Sync(func() {
			if wm.webview2InitFunctions[i].wv2.wv.Embed(uintptr(wm.webview2InitFunctions[i].wv2.iv.Handle())) {
				wm.webview2InitFunctions[i].wv2.inited = true
				if wm.webview2InitFunctions[i].succssFunc != nil {
					wm.webview2InitFunctions[i].succssFunc()
				}
				wm.webview2InitFunctions[i].wv2.wv.Resize()
			} else {
				if wm.webview2InitFunctions[i].failFunc != nil {
					wm.webview2InitFunctions[i].failFunc()
				}
			}
		})
	}
}

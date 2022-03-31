package walkmgr

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pirogom/walk"
	"github.com/pirogom/win"
)

var (
	defaultIcon    *walk.Icon
	defaultWinName string = "default walkmgr winname"
)

/**
*	MsgBox
**/
func MsgBox(msg string, window ...*walk.MainWindow) {
	if len(window) > 0 {
		walk.MsgBox(window[0], "알림", msg, walk.MsgBoxOK|walk.MsgBoxSetForeground)
	} else {
		walk.MsgBox(nil, "알림", msg, walk.MsgBoxOK|walk.MsgBoxSetForeground)
	}
}

/**
*	Confirm
**/
func Confirm(msg string, window ...*walk.MainWindow) bool {
	if len(window) > 0 {
		if walk.MsgBox(window[0], "알림", msg, walk.MsgBoxYesNo|walk.MsgBoxSetForeground) == win.IDNO {
			return false
		} else {
			return true
		}
	} else {
		if walk.MsgBox(nil, "알림", msg, walk.MsgBoxYesNo|walk.MsgBoxSetForeground) == win.IDNO {
			return false
		} else {
			return true
		}
	}
}

/**
*	LoadIcon
**/
func LoadIcon(icoBuf []byte, icoName string) {
	icoFile := filepath.Join(os.TempDir(), icoName)

	var err error

	if _, err = os.Stat(icoName); os.IsNotExist(err) {
		if err = ioutil.WriteFile(icoFile, icoBuf, 0644); err != nil {
			return
		}
	}
	defaultIcon, _ = walk.NewIconFromFile(icoFile)
}

/**
*	LoadIconFromFile
**/
func LoadIconFromFile(icoPath string) {
	defaultIcon, _ = walk.NewIconFromFile(icoPath)
}

/**
*	SetDefaultWindowName
**/
func SetDefaultWindowName(name string) {
	defaultWinName = name
}

/**
*	SetMultiLineText
**/
func MultiLineText(text string, limit int) string {
	runeText := []rune(text)
	var res string

	for {
		if len(runeText) <= limit {
			res += string(runeText)
			break
		}

		res += string(runeText[:limit]) + "\n"
		runeText = runeText[limit:]
	}
	return res
}

/**
*	CenterPos
**/
func CenterPos(width int, height int) (int, int) {
	var x, y int
	var rtDesk win.RECT
	win.GetWindowRect(win.GetDesktopWindow(), &rtDesk)

	x = (int(rtDesk.Right) - width) / 2
	y = (int(rtDesk.Bottom) - height) / 2

	return x, y
}

/**
*	AdsPos
**/
func AdsPos(width int, height int) (int, int) {
	var x, y int
	var rtDesk win.RECT
	win.GetWindowRect(win.GetDesktopWindow(), &rtDesk)

	x = int(rtDesk.Right) - width
	y = int(rtDesk.Bottom) - (height + 40)

	return x, y
}

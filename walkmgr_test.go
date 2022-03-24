package walkmgr

import (
	"testing"

	"github.com/pirogom/walk"
)

func TestDefaultWin(t *testing.T) {
	wm := NewWin("기본윈도", 640, 480)
	wm.Start()
}

func TestForegroundWin(t *testing.T) {
	wm := NewWin("기본윈도-최상단에 생성", 640, 480)
	wm.StartForeground()
}

func TestNoMinBox(t *testing.T) {
	wm := NewWin("기본윈도-최소화X", 640, 480)
	wm.DisableMinBox()

	wm.Start()
}

func TestNoMaxBox(t *testing.T) {
	wm := NewWin("기본윈도-최대화X", 640, 480)
	wm.DisableMaxBox()

	wm.Start()
}

func TestNoSysMenu(t *testing.T) {
	wm := NewWin("기본윈도-시스템메뉴X", 640, 480)
	wm.DisableSysmenu()

	wm.Start()
}

func TestNoTitle(t *testing.T) {
	wm := NewWin("기본윈도-타이틀바X", 640, 480)
	wm.DisableTitleBar()

	wm.Start()
}

func TestNoResize(t *testing.T) {
	wm := NewWin("기본윈도-사이즈변경X", 640, 480)
	wm.NoResize()
	wm.Start()
}

func TestUI(t *testing.T) {
	wm := NewWin("UI테스트", 640, 1080)

	// 버튼
	wm.PushButton("버튼", func() {

	})

	// 체크박스
	wm.CheckBox("체크박스(체크)", true, func() {

	})
	wm.CheckBox("체크박스(체크X)", false, func() {

	})

	// 드롭다운
	wm.DropDownBox([]string{"하나", "둘", "셋", "넷"})
	wm.DropDownBox([]string{"다섯", "여섯", "일곱", "여덟"}, 2)

	// 에디트 박스
	wm.NumberEdit()
	wm.NumberEdit(100)

	wm.LineEdit()
	wm.LineEdit("LineEdit 수정가능")

	wm.LineStatic()
	wm.LineStatic("LineStatic 수정불가")

	wm.TextEdit()
	wm.TextEdit("TextEdit 수정가능")

	wm.TextStatic()
	wm.TextStatic("TextStatic 수정불가")

	wm.TextArea()
	wm.TextArea("TextArea 수정가능")

	wm.TextAreaStatic()
	wm.TextAreaStatic("TextAreaStatic 수정불가")

	wm.Label("라벨 - 기본")
	wm.Label("라벨 - LEFT", ALIGN_LEFT)
	wm.Label("라벨 - CENTER", ALIGN_CENTER)
	wm.Label("라벨 - RIGHT", ALIGN_RIGHT)

	wm.Slider(0, 100, 50)

	wm.Start()
}

func TestImageView(t *testing.T) {
	img := LoadImage(".\\test_data\\gopher.png")

	if img == nil {
		return
	}

	defer func(img *walk.Image) {
		if img != nil {
			(*img).Dispose()
			img = nil
		}
	}(img)

	wm := NewWin("이미지뷰", 640, 480)

	iv := wm.ImageView(IV_ZOOM)
	iv.SetImage(*img)

	wm.Start()
}

func TestWebView(t *testing.T) {
	wm := NewWin("웹뷰(Alert Disabled)", 640, 480)

	wm.WebView("https://modu-print.tistory.com")

	wm.Start()
}

func TestWebViewWithAlert(t *testing.T) {
	wm := NewWin("웹뷰(Alert Enabled)", 640, 480)

	wm.WebViewWithAlert("https://modu-print.tistory.com")

	wm.Start()
}

func TestIgnoreClose(t *testing.T) {
	wm := NewWin("닫지 못하는 창", 640, 480)

	wm.PushButton("강제닫음", func() {
		wm.ForceClose()
	})
	wm.IgnoreClosing()

	wm.Start()
}

func TestVertLayout(t *testing.T) {
	//wm := NewWin("", 640, 480) default layout is vertical
	wm := NewWin("LAYOUT_VERT", 640, 480, LAYOUT_VERT)

	wm.PushButton("버튼1", func() {
	})
	wm.PushButton("버튼2", func() {
	})

	wm.Start()
}

func TestHoriLayout(t *testing.T) {
	wm := NewWin("LAYOUT_HORI", 640, 480, LAYOUT_HORI)

	wm.PushButton("버튼1", func() {
	})
	wm.PushButton("버튼2", func() {
	})

	wm.Start()
}

func TestHSplit(t *testing.T) {
	wm := NewWin("HSplit", 640, 480)

	wm.Split(LAYOUT_HORI)
	wm.PushButton("버튼1", func() {
	})
	wm.PushButton("버튼2", func() {
	})
	wm.End()

	wm.Start()
}

func TestVSplit(t *testing.T) {
	wm := NewWin("VSplit", 640, 480)

	wm.Split(LAYOUT_VERT)
	wm.PushButton("버튼1", func() {
	})
	wm.PushButton("버튼2", func() {
	})
	wm.End()

	wm.Start()
}

func TestGroupBox(t *testing.T) {
	wm := NewWin("GroupBox(vert)", 640, 480)

	wm.GroupBox("그룹박스(vert)", LAYOUT_VERT)
	wm.PushButton("버튼1", func() {
	})
	wm.PushButton("버튼2", func() {
	})
	wm.End()

	wm.Start()
}

func TestGroupBox2(t *testing.T) {
	wm := NewWin("GroupBox(hori)", 640, 480)

	wm.GroupBox("그룹박스(hori)", LAYOUT_HORI)
	wm.PushButton("버튼1", func() {
	})
	wm.PushButton("버튼2", func() {
	})
	wm.End()

	wm.Start()
}

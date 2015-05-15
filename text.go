package termui

const dot = '…'

type Text struct {
	BaseElement

	txt       []rune
	renderTxt []rune
}

var _ Element = new(Text) // Interface checking...

func NewText(txt string) *Text {
	ft := new(Text)
	ft.SetText(txt)
	return ft
}

func (t *Text) Width() int {
	return len(t.renderTxt)
}

func (t *Text) Height() int {
	return 1
}

func (t *Text) Children() []Element {
	return []Element{}
}

func (t *Text) Measure(availableWidth, availableHeight int) (width int, height int) {
	width = len(t.txt)
	if width > availableWidth && availableWidth > 0 {
		width = availableWidth
	}

	height = 1
	if height > availableHeight && availableHeight > 0 {
		height = availableHeight
	}

	return
}

func (t *Text) Arrange(finalWidth, finalHeight int) {
	tLen := len(t.txt)
	if finalWidth < tLen {
		t.renderTxt = append(t.txt[:finalWidth-1], dot)
	} else {
		t.renderTxt = t.txt
	}
}

func (t *Text) Render(rn Renderer) {
	for i, r := range t.renderTxt {
		rn.Set(i, 0, r)
	}
}

func (t *Text) SetText(txt string) {
	t.txt = []rune(txt)
}

func (t *Text) Text() string {
	return string(t.txt)
}

func (t *Text) Name() string {
	return "text"
}

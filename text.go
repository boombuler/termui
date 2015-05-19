package termui

import (
	"github.com/boombuler/termui/css"
)

const dot = '…'

// Text is an element which simply draws a static text.
type Text struct {
	BaseElement

	txt       []rune
	renderTxt []rune
}

var _ Element = new(Text) // Interface checking...

// NewText creates a new Text element.
func NewText(txt string) *Text {
	ft := new(Text)
	ft.SetText(txt)
	return ft
}

// Width returns the width of the text
func (t *Text) Width() int {
	return len(t.renderTxt)
}

// Height returns the height of the text.
func (t *Text) Height() int {
	return 1
}

// Children returns all nested elements.
func (t *Text) Children() []css.Styleable {
	return nil
}

// Measure gets the "wanted" size of the element based on the available size
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

// Arrange sets the final size for the Element end tells it to Arrange itself
func (t *Text) Arrange(finalWidth, finalHeight int) {
	tLen := len(t.txt)
	if finalWidth < tLen {
		t.renderTxt = append(t.txt[:finalWidth-1], dot)
	} else {
		t.renderTxt = t.txt
	}
}

// Render renders the element on the given Renderer
func (t *Text) Render(rn Renderer) {
	for i, r := range t.renderTxt {
		rn.Set(i, 0, r)
	}
}

// SetText replaces the the text of the text element.
func (t *Text) SetText(txt string) {
	t.txt = []rune(txt)
}

// Text returns the current text of the element.
func (t *Text) Text() string {
	return string(t.txt)
}

// Name returns the constant name of the text element for css styling.
func (t *Text) Name() string {
	return "text"
}

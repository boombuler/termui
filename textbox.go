package termui

import (
	"github.com/nsf/termbox-go"
	"time"
	"unicode/utf8"
)

type TextBox struct {
	BaseElement
	width int

	ticker         *time.Ticker
	blinkOn        bool
	text           []byte
	line_voffset   int
	cursor_boffset int // cursor offset in bytes
	cursor_coffset int // cursor offset in unicode code points
}

const blink_freq = 500 * time.Millisecond

var _ FocusElement = &TextBox{}

func NewTextBox() *TextBox {
	return new(TextBox)
}

func (tb *TextBox) Name() string {
	return "textbox"
}

func (tb *TextBox) Measure(availableWidth, availableHeight int) (width int, height int) {
	return availableWidth, 1
}

func (tb *TextBox) Arrange(finalWidth, finalHeight int) {
	const preferred_horizontal_threshold = 5

	tb.width = finalWidth

	ht := preferred_horizontal_threshold
	max_h_threshold := (tb.width - 1) / 2
	if ht > max_h_threshold {
		ht = max_h_threshold
	}

	threshold := tb.width - 1
	if tb.line_voffset != 0 {
		threshold = tb.width - ht
	}
	if tb.cursor_coffset-tb.line_voffset >= threshold {
		tb.line_voffset = tb.cursor_coffset + (ht - tb.width + 1)
	}

	if tb.line_voffset != 0 && tb.cursor_coffset-tb.line_voffset < ht {
		tb.line_voffset = tb.cursor_coffset - ht
		if tb.line_voffset < 0 {
			tb.line_voffset = 0
		}
	}
}

func (tb *TextBox) Children() []Element {
	return []Element{}
}

func (tb *TextBox) Height() int {
	return 1
}
func (tb *TextBox) Width() int {
	return tb.width
}

func (tb *TextBox) SetFocused(v bool) {
	if v {
		ticker := time.NewTicker(blink_freq)
		tb.ticker = ticker
		go func() {
			for _ = range ticker.C {
				tb.blinkOn = !tb.blinkOn
				Update()
			}
		}()
	} else {
		tb.ticker.Stop()
		tb.ticker = nil
		tb.blinkOn = false
	}
}

func (tb *TextBox) Render(fn Renderer) {
	cursorPos := tb.cursor_coffset - tb.line_voffset

	for lx := 0; lx < tb.width; lx++ {
		if lx == cursorPos && tb.blinkOn {
			fn.SetAttr(lx, 0, ' ', termbox.AttrReverse)
		} else {
			fn.Set(lx, 0, ' ')
		}

	}

	t := tb.text
	lx := 0
	for {
		rx := lx - tb.line_voffset
		if len(t) == 0 {
			break
		}

		if rx >= tb.width {
			fn.Set(tb.width-1, 0, '→')
			break
		}

		r, size := utf8.DecodeRune(t)
		if rx >= 0 {
			if rx == cursorPos && tb.blinkOn {
				fn.SetAttr(rx, 0, r, termbox.AttrReverse)
			} else {
				fn.Set(rx, 0, r)
			}
		}
		lx += 1
		t = t[size:]
	}

	if tb.line_voffset != 0 {
		fn.Set(0, 0, '←')
	}
}

func (tb *TextBox) HandleKey(k termbox.Key, ch rune) bool {
	switch k {
	case termbox.KeyArrowLeft, termbox.KeyCtrlB:
		tb.MoveCursorOneRuneBackward()
	case termbox.KeyArrowRight, termbox.KeyCtrlF:
		tb.MoveCursorOneRuneForward()
	case termbox.KeyBackspace, termbox.KeyBackspace2:
		tb.DeleteRuneBackward()
	case termbox.KeyDelete, termbox.KeyCtrlD:
		tb.DeleteRuneForward()
	case termbox.KeySpace:
		tb.InsertRune(' ')
	case termbox.KeyCtrlK:
		tb.DeleteTheRestOfTheLine()
	case termbox.KeyHome, termbox.KeyCtrlA:
		tb.MoveCursorToBeginningOfTheLine()
	case termbox.KeyEnd, termbox.KeyCtrlE:
		tb.MoveCursorToEndOfTheLine()
	default:
		if ch != 0 {
			tb.InsertRune(ch)
		}
		return false
	}
	return true
}

func (tb *TextBox) moveCursorTo(boffset int) {
	tb.cursor_boffset = boffset
	text := tb.text[:boffset]
	coffset := 0
	for len(text) > 0 {
		_, size := utf8.DecodeRune(text)
		text = text[size:]
		coffset += 1
	}
	tb.cursor_coffset = coffset
}

func (tb *TextBox) runeUnderCursor() (rune, int) {
	return utf8.DecodeRune(tb.text[tb.cursor_boffset:])
}

func (tb *TextBox) runeBeforeCursor() (rune, int) {
	return utf8.DecodeLastRune(tb.text[:tb.cursor_boffset])
}

func (tb *TextBox) MoveCursorOneRuneBackward() {
	if tb.cursor_boffset == 0 {
		return
	}
	_, size := tb.runeBeforeCursor()
	tb.moveCursorTo(tb.cursor_boffset - size)
}

func (tb *TextBox) MoveCursorOneRuneForward() {
	if tb.cursor_boffset == len(tb.text) {
		return
	}
	_, size := tb.runeUnderCursor()
	tb.moveCursorTo(tb.cursor_boffset + size)
}

func (tb *TextBox) MoveCursorToBeginningOfTheLine() {
	tb.moveCursorTo(0)
}

func (tb *TextBox) MoveCursorToEndOfTheLine() {
	tb.moveCursorTo(len(tb.text))
}

func (tb *TextBox) DeleteRuneBackward() {
	if tb.cursor_boffset == 0 {
		return
	}

	tb.MoveCursorOneRuneBackward()
	_, size := tb.runeUnderCursor()
	tb.text = byte_slice_remove(tb.text, tb.cursor_boffset, tb.cursor_boffset+size)
}

func (tb *TextBox) DeleteRuneForward() {
	if tb.cursor_boffset == len(tb.text) {
		return
	}
	_, size := tb.runeUnderCursor()
	tb.text = byte_slice_remove(tb.text, tb.cursor_boffset, tb.cursor_boffset+size)
}

func (tb *TextBox) DeleteTheRestOfTheLine() {
	tb.text = tb.text[:tb.cursor_boffset]
}

func (tb *TextBox) InsertRune(r rune) {
	var buf [utf8.UTFMax]byte
	n := utf8.EncodeRune(buf[:], r)
	tb.text = byte_slice_insert(tb.text, tb.cursor_boffset, buf[:n])
	tb.MoveCursorOneRuneForward()
}

func byte_slice_grow(s []byte, desired_cap int) []byte {
	if cap(s) < desired_cap {
		ns := make([]byte, len(s), desired_cap)
		copy(ns, s)
		return ns
	}
	return s
}

func byte_slice_remove(text []byte, from, to int) []byte {
	size := to - from
	copy(text[from:], text[to:])
	text = text[:len(text)-size]
	return text
}

func byte_slice_insert(text []byte, offset int, what []byte) []byte {
	n := len(text) + len(what)
	text = byte_slice_grow(text, n)
	text = text[:n]
	copy(text[offset+len(what):], text[offset:])
	copy(text[offset:], what)
	return text
}

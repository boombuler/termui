package termui

import (
	"github.com/nsf/termbox-go"
	"time"
	"unicode/utf8"
)

// TextBox is an UI element which can be used to query text input from the user.
type TextBox struct {
	BaseElement
	width int

	stopTicker  chan struct{}
	ticker      *time.Ticker
	blinkOn     bool
	text        []byte
	lineVOffset int
	// cursor offset in bytes
	cursorBOffset int
	// cursor offset in unicode code points
	cursorCOoffset int
}

const blinkFreq = 500 * time.Millisecond

var _ FocusElement = &TextBox{}

// NewTextBox creates a new textbox element.
func NewTextBox() *TextBox {
	return new(TextBox)
}

// Name returns the constant name of the textbox for css styling.
func (tb *TextBox) Name() string {
	return "textbox"
}

// Measure gets the "wanted" size of the element based on the available size
func (tb *TextBox) Measure(availableWidth, availableHeight int) (width int, height int) {
	width = utf8.RuneCount(tb.text)
	if width > availableWidth && availableWidth > 0 {
		width = availableWidth
	}

	height = 1
	if height > availableHeight && availableHeight > 0 {
		height = availableHeight
	}
	return
}

// Text returns the current text of the textbox
func (tb *TextBox) Text() string {
	return string(tb.text)
}

// SetText sets the text of the textbox and moves the cursor to the end.
func (tb *TextBox) SetText(s string) {
	tb.text = []byte(s)
	tb.MoveCursorToEndOfTheLine()
}

// Arrange sets the final size for the Element end tells it to Arrange itself
func (tb *TextBox) Arrange(finalWidth, finalHeight int) {
	const preferredHorizontalThreshold = 5

	tb.width = finalWidth

	ht := preferredHorizontalThreshold
	maxHThreshold := (tb.width - 1) / 2
	if ht > maxHThreshold {
		ht = maxHThreshold
	}

	threshold := tb.width - 1
	if tb.lineVOffset != 0 {
		threshold = tb.width - ht
	}
	if tb.cursorCOoffset-tb.lineVOffset >= threshold {
		tb.lineVOffset = tb.cursorCOoffset + (ht - tb.width + 1)
	}

	if tb.lineVOffset != 0 && tb.cursorCOoffset-tb.lineVOffset < ht {
		tb.lineVOffset = tb.cursorCOoffset - ht
		if tb.lineVOffset < 0 {
			tb.lineVOffset = 0
		}
	}
}

// Children returns all nested elements.
func (tb *TextBox) Children() []Element {
	return []Element{}
}

// Height returns the height of the textbox.
func (tb *TextBox) Height() int {
	return 1
}

// Width returns the width of the textbox
func (tb *TextBox) Width() int {
	return tb.width
}

// SetFocused is called by the ui system to indicate that the element has the focus.
func (tb *TextBox) SetFocused(v bool) {
	if v {
		ticker := time.NewTicker(blinkFreq)
		tb.ticker = ticker
		tb.stopTicker = make(chan struct{})
		tb.blinkOn = true
		go func() {
			for {
				select {
				case <-ticker.C:
					tb.blinkOn = !tb.blinkOn
					Update()
				case <-tb.stopTicker:
					tb.blinkOn = false
					Update()
					return
				}
			}
		}()
	} else {
		tb.ticker.Stop()
		tb.ticker = nil
		tb.stopTicker <- struct{}{}
		close(tb.stopTicker)
	}
}

// Render renders the element on the given Renderer
func (tb *TextBox) Render(fn Renderer) {
	cursorPos := tb.cursorCOoffset - tb.lineVOffset

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
		rx := lx - tb.lineVOffset
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
		lx++
		t = t[size:]
	}

	if tb.lineVOffset != 0 {
		fn.Set(0, 0, '←')
	}
}

// HandleKey is called by the ui system if the element is asked to process a key.
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
	tb.cursorBOffset = boffset
	text := tb.text[:boffset]
	coffset := 0
	for len(text) > 0 {
		_, size := utf8.DecodeRune(text)
		text = text[size:]
		coffset++
	}
	tb.cursorCOoffset = coffset
}

func (tb *TextBox) runeUnderCursor() (rune, int) {
	return utf8.DecodeRune(tb.text[tb.cursorBOffset:])
}

func (tb *TextBox) runeBeforeCursor() (rune, int) {
	return utf8.DecodeLastRune(tb.text[:tb.cursorBOffset])
}

// MoveCursorOneRuneBackward moves the cursor one rune to the left.
func (tb *TextBox) MoveCursorOneRuneBackward() {
	if tb.cursorBOffset == 0 {
		return
	}
	_, size := tb.runeBeforeCursor()
	tb.moveCursorTo(tb.cursorBOffset - size)
}

// MoveCursorOneRuneForward moves the cursor one rune to the right
func (tb *TextBox) MoveCursorOneRuneForward() {
	if tb.cursorBOffset == len(tb.text) {
		return
	}
	_, size := tb.runeUnderCursor()
	tb.moveCursorTo(tb.cursorBOffset + size)
}

// MoveCursorToBeginningOfTheLine moves the cursor to the beginning of the text.
func (tb *TextBox) MoveCursorToBeginningOfTheLine() {
	tb.moveCursorTo(0)
}

// MoveCursorToEndOfTheLine moves the cursor to the end of the line
func (tb *TextBox) MoveCursorToEndOfTheLine() {
	tb.moveCursorTo(len(tb.text))
}

// DeleteRuneBackward deletes the rune before the current cursor position
func (tb *TextBox) DeleteRuneBackward() {
	if tb.cursorBOffset == 0 {
		return
	}

	tb.MoveCursorOneRuneBackward()
	_, size := tb.runeUnderCursor()
	tb.text = byteSliceRemove(tb.text, tb.cursorBOffset, tb.cursorBOffset+size)
}

// DeleteRuneForward deletes the rune which is at the current cursor position
func (tb *TextBox) DeleteRuneForward() {
	if tb.cursorBOffset == len(tb.text) {
		return
	}
	_, size := tb.runeUnderCursor()
	tb.text = byteSliceRemove(tb.text, tb.cursorBOffset, tb.cursorBOffset+size)
}

// DeleteTheRestOfTheLine removes all text from the current position to the end of the text.
func (tb *TextBox) DeleteTheRestOfTheLine() {
	tb.text = tb.text[:tb.cursorBOffset]
}

// InsertRune inserts the rune at the current cursor position
func (tb *TextBox) InsertRune(r rune) {
	var buf [utf8.UTFMax]byte
	n := utf8.EncodeRune(buf[:], r)
	tb.text = byteSliceInsert(tb.text, tb.cursorBOffset, buf[:n])
	tb.MoveCursorOneRuneForward()
}

func byteSliceGrow(s []byte, desiredCap int) []byte {
	if cap(s) < desiredCap {
		ns := make([]byte, len(s), desiredCap)
		copy(ns, s)
		return ns
	}
	return s
}

func byteSliceRemove(text []byte, from, to int) []byte {
	size := to - from
	copy(text[from:], text[to:])
	text = text[:len(text)-size]
	return text
}

func byteSliceInsert(text []byte, offset int, what []byte) []byte {
	n := len(text) + len(what)
	text = byteSliceGrow(text, n)
	text = text[:n]
	copy(text[offset+len(what):], text[offset:])
	copy(text[offset:], what)
	return text
}

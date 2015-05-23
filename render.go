package termui

import (
	"github.com/nsf/termbox-go"
)

var rootrenderer = Renderer{0, 0, nil, termbox.ColorDefault, termbox.ColorDefault}

// Renderer helps drawing an element
type Renderer struct {
	x, y   int
	cur    Element
	fg, bg termbox.Attribute
}

// Set the given rune at a relative x and y position
func (r Renderer) Set(x, y int, ch rune) {
	termbox.SetCell(x+r.x, y+r.y, ch, r.fg, r.bg)
}

// SetAttr sets the given rune at a relative x and y position and applys the attribute to the foreground.
// This can be used to invert or underline a cell.
func (r Renderer) SetAttr(x, y int, ch rune, attr termbox.Attribute) {
	termbox.SetCell(x+r.x, y+r.y, ch, r.fg|attr, r.bg)
}

// RenderChild draws the given child.
func (r Renderer) RenderChild(e Element, width, height, xOffset, yOffset int) {
	elW, elH := e.Width(), e.Height()

	gravity := GravityProperty.Get(e)
	fg := ForegroundProperty.Get(e)
	bg := BackgroundProperty.Get(e)

	margin := MarginProperty.Get(e)
	width -= margin.Right + margin.Left
	height -= margin.Bottom + margin.Top
	xOffset += margin.Left
	yOffset += margin.Top

	var spaceLeft int
	switch gravity & horizontal {
	case Left, Left | Right:
		spaceLeft = 0
	case Right:
		spaceLeft = width - elW
	case NoGravity:
		spaceLeft = (width - elW) / 2
	}
	xOffset += spaceLeft
	width -= spaceLeft

	var spaceTop int
	switch gravity & vertical {
	case Top, Top | Bottom:
		spaceTop = 0
	case Bottom:
		spaceTop = height - elH
	case NoGravity:
		spaceTop = (height - elH) / 2
	}
	yOffset += spaceTop
	height -= spaceTop

	if r.cur != nil {
		for x := xOffset; x < xOffset+width; x++ {
			for y := yOffset; y < yOffset+height; y++ {
				termbox.SetCell(x+r.x, y+r.y, ' ', fg, bg)
			}
		}
	}

	e.Render(Renderer{
		x:  r.x + xOffset,
		y:  r.y + yOffset,
		fg: fg, bg: bg,
		cur: e,
	})
}

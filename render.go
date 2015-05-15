package termui

import (
	"github.com/boombuler/termui/css"
	"github.com/nsf/termbox-go"
)

var rootrenderer Renderer = Renderer{0, 0, nil, termbox.ColorDefault, termbox.ColorDefault}

type Renderer struct {
	x, y   int
	cur    Element
	fg, bg termbox.Attribute
}

func (r Renderer) Set(x, y int, ch rune) {
	termbox.SetCell(x+r.x, y+r.y, ch, r.fg, r.bg)
}

func (r Renderer) SetAttr(x, y int, ch rune, attr termbox.Attribute) {
	termbox.SetCell(x+r.x, y+r.y, ch, r.fg|attr, r.bg)
}

func (r Renderer) RenderChild(e Element, width, height, xOffset, yOffset int) {
	elW, elH := e.Width(), e.Height()

	style := css.Get(e)

	var ok bool
	var gravity Gravity
	if gravity, ok = style.Value(GravityProperty).(Gravity); !ok {
		gravity = NoGravity
	}

	var fg, bg termbox.Attribute
	if fg, ok = style.Value(ForegroundProperty).(termbox.Attribute); !ok {
		fg = termbox.ColorDefault
	}
	if bg, ok = style.Value(BackgroundProperty).(termbox.Attribute); !ok {
		bg = termbox.ColorDefault
	}

	switch gravity & horizontal {
	case Left, Left | Right:
		xOffset += 0
	case Right:
		xOffset += width - elW
	case NoGravity:
		xOffset += (width - elW) / 2
	}
	switch gravity & vertical {
	case Top, Top | Bottom:
		yOffset += 0
	case Bottom:
		yOffset += height - elH
	case NoGravity:
		yOffset += (height - elH) / 2
	}

	e.Render(Renderer{
		x:  r.x + xOffset,
		y:  r.y + yOffset,
		fg: fg, bg: bg,
	})
}

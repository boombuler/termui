package termui

import (
	"github.com/boombuler/termui/css"
	"github.com/nsf/termbox-go"
)

type Gravity byte

const (
	NoGravity Gravity = 0
	Top       Gravity = 1 << iota
	Left
	Right
	Bottom

	horizontal = Left | Right
	vertical   = Top | Bottom
)

type Element interface {
	css.Styleable
	Measure(availableWidth, availableHeight int) (width int, height int)
	Arrange(finalWidth, finalHeight int)
	Render(fn Renderer)
	Width() int
	Height() int

	Children() []Element
	SetParent(e Element)
}

type FocusElement interface {
	Element
	SetFocused(v bool)
	HandleKey(k termbox.Key, ch rune) bool
}

type BaseElement struct {
	css.IdAndClasses
	parent Element
}

func (be *BaseElement) SetParent(e Element) {
	be.parent = e
}

func (be *BaseElement) Parent() css.Styleable {
	return be.parent
}

func getGravity(e Element) Gravity {
	style := css.Get(e)
	gravity := NoGravity
	if gravAny, ok := style[GravityProperty]; ok {
		if gravity, ok = gravAny.(Gravity); !ok {
			gravity = NoGravity
		}
	}
	return gravity
}

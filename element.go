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
	// Measure gets the "wanted" size of the element based on the available size
	Measure(availableWidth, availableHeight int) (width int, height int)
	// Arrange sets the final size for the Element end tells it to Arrange itself
	Arrange(finalWidth, finalHeight int)
	// Render renders the element on the given Renderer
	Render(fn Renderer)
	// Width returns the width of the border
	Width() int
	// Height returns the height of the element
	Height() int

	// Children returns all nested child elements of the element
	Children() []Element
	// SetParent sets the parent element of this element
	SetParent(e Element)
}

type FocusElement interface {
	Element
	SetFocused(v bool)
	HandleKey(k termbox.Key, ch rune) bool
}

type BaseElement struct {
	css.IDAndClasses
	parent Element
}

// SetParent sets the parent element of this element
func (be *BaseElement) SetParent(e Element) {
	be.parent = e
}

// Parent returns the parent of the element
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

package termui

import (
	"github.com/boombuler/termui/css"
	"github.com/nsf/termbox-go"
)

// Gravity defines the horizontal and vertical orientation of an element.
type Gravity byte

const (
	// NoGravity -> element is centered horizontal and vertical
	NoGravity Gravity = 0
	// Top pulls the element to the top
	Top Gravity = 1 << iota
	// Left pulls the element to the left
	Left
	// Right pulls the element to the right
	Right
	// Bottom pulls the element to the top bottom
	Bottom

	horizontal = Left | Right
	vertical   = Top | Bottom
)

// Element must be implemented by any UI element.
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

	// SetParent sets the parent element of this element
	SetParent(e Element)
}

type KeyHandler interface {
	Element
	// HandleKey is called by the ui system if the element is asked to process a key. It should return true if the key
	// was successfully handled and should not be processed by other elements.
	HandleKey(k termbox.Key, ch rune) bool
}

// FocusElement is an element which can get the input focus.
type FocusElement interface {
	KeyHandler
	// SetFocused is called by the ui system to indicate that the element has the focus.
	SetFocused(v bool)
}

// BaseElement helps implementing the Element interface for ui elements. It handles classes, the ID and also the parent element.
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

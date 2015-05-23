package termui

import (
	"github.com/boombuler/termui/css"
)

type WrapPanel struct {
	BaseElement
	children *ElementCollection

	width, height int
	orientation   Orientation
	childpos      map[Element]rect
}

type rect struct {
	x, y, width, height int
}

var _ Element = new(WrapPanel)

func NewWrapPanel(o Orientation) *WrapPanel {
	wp := &WrapPanel{
		orientation: o,
		children:    new(ElementCollection),
		childpos:    make(map[Element]rect),
	}
	return wp
}

func (v *WrapPanel) Add(e ...Element) {
	v.children.Add(e...)
	for _, el := range e {
		el.SetParent(v)
	}
}

func (v *WrapPanel) Remove(e Element) {
	idx := v.children.IndexOf(e)
	if idx != -1 {
		e.SetParent(nil)
		v.children.RemoveAt(idx)
	}
}

func (v *WrapPanel) Clear() {
	for _, el := range v.children.Items() {
		el.SetParent(nil)
		delete(v.childpos, el)
	}

	v.children.Clear()
}

func (v *WrapPanel) Len() int {
	return v.children.Len()
}

// Name returns the constant name of the WrapPanel for css styling.
func (v *WrapPanel) Name() string {
	return "wrappanel"
}

// Children returns all nested elements.
func (v *WrapPanel) Children() []css.Styleable {
	res := make([]css.Styleable, v.Len())
	for i := 0; i < v.Len(); i++ {
		res[i] = v.children.At(i)
	}
	return res
}

func (v *WrapPanel) measureVertical(availableWidth, availableHeight int, arrange bool) (width int, height int) {
	cIdx := 0
	width = 0
	height = 0
	for cIdx < v.Len() {
		colWidth := 0
		colHeight := 0
		cStart := cIdx
		cEnd := cIdx

		for i := cIdx; i < v.Len(); i++ {
			child := v.children.At(i)
			cw, ch := MeasureChild(child, 0, 0)

			if colHeight+ch > availableHeight && availableHeight != 0 {
				if colHeight > height {
					height = colHeight
				}
				break
			}

			if arrange {
				v.childpos[child] = rect{
					x:      width,
					width:  cw,
					height: ch,
				}
			}

			if cw > colWidth {
				colWidth = cw
			}
			colHeight += ch
			cEnd = cIdx
			cIdx++
		}
		width += colWidth

		if arrange {
			y := 0
			for i := cStart; i <= cEnd; i++ {
				child := v.children.At(i)
				rect := v.childpos[child]
				ArrangeChild(child, rect.width, rect.height)
				rect.y = y
				y += rect.height
				v.childpos[child] = rect
			}
		}
	}
	return
}

func (v *WrapPanel) measureHorizontal(availableWidth, availableHeight int, arrange bool) (width int, height int) {
	cIdx := 0
	width = 0
	height = 0
	for cIdx < v.Len() {
		colWidth := 0
		colHeight := 0
		cStart := cIdx
		cEnd := cIdx

		for i := cIdx; i < v.Len(); i++ {
			child := v.children.At(i)
			cw, ch := MeasureChild(child, 0, 0)

			if colWidth+cw > availableWidth && availableWidth != 0 {
				if colWidth > width {
					width = colWidth
				}
				break
			}

			if arrange {
				v.childpos[child] = rect{
					y:      height,
					width:  cw,
					height: ch,
				}
			}

			if ch > colHeight {
				colHeight = ch
			}
			colWidth += cw
			cEnd = cIdx
			cIdx++
		}
		height += colHeight

		if arrange {
			x := 0
			for i := cStart; i <= cEnd; i++ {
				child := v.children.At(i)
				rect := v.childpos[child]
				ArrangeChild(child, rect.width, rect.height)
				rect.x = x
				x += rect.width
				v.childpos[child] = rect
			}
		}
	}
	return
}

// Measure gets the "wanted" size of the element based on the available size
func (v *WrapPanel) Measure(availableWidth, availableHeight int) (width int, height int) {
	if v.orientation == Vertical {
		return v.measureVertical(availableWidth, availableHeight, false)
	} else {
		return v.measureHorizontal(availableWidth, availableHeight, false)
	}
}

// Arrange sets the final size for the Element end tells it to Arrange itself
func (v *WrapPanel) Arrange(finalWidth, finalHeight int) {
	v.width, v.height = finalWidth, finalHeight
	if v.orientation == Vertical {
		v.measureVertical(finalWidth, finalHeight, true)
	} else {
		v.measureHorizontal(finalWidth, finalHeight, true)
	}
}

// Render renders the element on the given Renderer
func (v *WrapPanel) Render(rn Renderer) {
	for el, pos := range v.childpos {
		rn.RenderChild(el, pos.width, pos.height, pos.x, pos.y)
	}
}

// Width returns the width of the WrapPanel
func (v *WrapPanel) Width() int {
	return v.width
}

// Height returns the height of the WrapPanel.
func (v *WrapPanel) Height() int {
	return v.height
}

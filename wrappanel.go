package termui

type WrapPanel struct {
	BaseElement

	width, height int
	orientation   Orientation
	children      []Element
	childpos      map[Element]rect
}

type rect struct {
	x, y, width, height int
}

var _ Element = new(WrapPanel)

func NewWrapPanel(o Orientation) *WrapPanel {
	return &WrapPanel{
		orientation: o,
		childpos:    make(map[Element]rect),
	}
}

// Name returns the constant name of the WrapPanel for css styling.
func (v *WrapPanel) Name() string {
	return "wrappanel"
}

// Children returns all nested elements.
func (v *WrapPanel) Children() []Element {
	return v.children
}

// AddChild adds the given children to the WrapPanel
func (v *WrapPanel) AddChild(e ...Element) {
	v.children = append(v.children, e...)
	for _, c := range e {
		c.SetParent(v)
	}
}

// Removes all children.
func (v *WrapPanel) Clear() {
	for _, c := range v.children {
		c.SetParent(nil)
	}
	v.childpos = make(map[Element]rect)
	v.children = nil
}

func (v *WrapPanel) measureVertical(availableWidth, availableHeight int, arrange bool) (width int, height int) {
	cIdx := 0
	width = 0
	height = 0
	for cIdx < len(v.children) {
		colWidth := 0
		colHeight := 0
		cStart := cIdx
		cEnd := cIdx

		for i := cIdx; i < len(v.children); i++ {
			cw, ch := v.children[i].Measure(0, 0)

			if colHeight+ch > availableHeight && availableHeight != 0 {
				if colHeight > height {
					height = colHeight
				}
				break
			}

			if arrange {
				v.childpos[v.children[i]] = rect{
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
				child := v.children[i]
				rect := v.childpos[child]
				child.Arrange(rect.width, rect.height)
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
	for cIdx < len(v.children) {
		colWidth := 0
		colHeight := 0
		cStart := cIdx
		cEnd := cIdx

		for i := cIdx; i < len(v.children); i++ {
			cw, ch := v.children[i].Measure(0, 0)

			if colWidth+cw > availableWidth && availableWidth != 0 {
				if colWidth > width {
					width = colWidth
				}
				break
			}

			if arrange {
				v.childpos[v.children[i]] = rect{
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
				child := v.children[i]
				rect := v.childpos[child]
				child.Arrange(rect.width, rect.height)
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

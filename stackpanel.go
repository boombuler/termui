package termui

type Orientation int

const (
	// Vertical orientation
	Vertical Orientation = iota
	// Horizontal orientation
	Horizontal
)

// StackPanel is an UI element which stacks multiple elements in vertical order.
type StackPanel struct {
	BaseElement

	orientation   Orientation
	childs        []Element
	heights       map[Element]int
	widths        map[Element]int
	width, height int
}

var _ Element = new(StackPanel) // Interface checking...

// NewStackPanel creates a new StackPanel
func NewStackPanel(o Orientation) *StackPanel {
	return &StackPanel{
		orientation: o,
	}
}

// Name returns the constant name of the StackPanel for css styling.
func (v *StackPanel) Name() string {
	return "stackpanel"
}

// Children returns all nested elements.
func (v *StackPanel) Children() []Element {
	return v.childs
}

// AddChild adds the given children to the StackPanel
func (v *StackPanel) AddChild(e ...Element) {
	v.childs = append(v.childs, e...)
	for _, c := range e {
		c.SetParent(v)
	}
}

// Measure gets the "wanted" size of the element based on the available size
func (v *StackPanel) Measure(availableWidth, availableHeight int) (width int, height int) {
	width, height = 0, 0
	v.heights = make(map[Element]int)
	v.widths = make(map[Element]int)
	for _, child := range v.childs {
		if v.orientation == Vertical {
			cw, ch := child.Measure(availableWidth, availableHeight-height)
			if cw > width {
				width = cw
			}
			height += ch
			v.heights[child] = ch
			v.widths[child] = cw
		} else {
			cw, ch := child.Measure(availableWidth-width, availableHeight)
			if ch > height {
				height = ch
			}
			width += cw
			v.heights[child] = ch
			v.widths[child] = cw
		}
	}
	return
}

// Arrange sets the final size for the Element end tells it to Arrange itself
func (v *StackPanel) Arrange(finalWidth, finalHeight int) {
	v.width, v.height = finalWidth, finalHeight
	v.Measure(finalWidth, finalHeight)

	for _, child := range v.childs {
		if v.orientation == Vertical {
			ch, ok := v.heights[child]
			if !ok {
				ch = 0
			}
			if finalHeight < ch {
				ch = finalHeight
			}
			cw, ok := v.widths[child]
			if !ok || cw > finalWidth {
				cw = finalWidth
			}

			child.Arrange(cw, ch)
			finalHeight -= ch
		} else {
			cw, ok := v.widths[child]
			if !ok {
				cw = 0
			}
			if finalWidth < cw {
				cw = finalWidth
			}
			ch, ok := v.heights[child]
			if !ok || ch > finalHeight {
				cw = finalHeight
			}

			child.Arrange(cw, ch)
			finalWidth -= cw
		}
	}
}

// Render renders the element on the given Renderer
func (v *StackPanel) Render(rn Renderer) {
	off := 0
	for _, child := range v.childs {
		if v.orientation == Vertical {
			if off >= v.height {
				break
			}
			height := child.Height()
			rn.RenderChild(child, v.width, height, 0, off)
			off += height
		} else {
			if off >= v.width {
				break
			}
			width := child.Width()
			rn.RenderChild(child, width, v.height, off, 0)
			off += width
		}
	}
}

// Width returns the width of the StackPanel
func (v *StackPanel) Width() int {
	if v.orientation == Vertical {
		return v.width
	}
	s := 0
	for _, child := range v.childs {
		s += child.Width()
	}
	return s
}

// Height returns the height of the StackPanel.
func (v *StackPanel) Height() int {
	if v.orientation == Horizontal {
		return v.height
	}
	s := 0
	for _, child := range v.childs {
		s += child.Height()
	}
	return s
}

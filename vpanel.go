package termui

// VPanel is an UI element which stacks multiple elements in vertical order.
type VPanel struct {
	BaseElement
	childs        []Element
	heights       map[Element]int
	widths        map[Element]int
	width, height int
}

var _ Element = new(VPanel) // Interface checking...

// NewVPanel creates a new VPanel
func NewVPanel() *VPanel {
	return new(VPanel)
}

// Name returns the constant name of the vpanel for css styling.
func (v *VPanel) Name() string {
	return "vpanel"
}

// Children returns all nested elements.
func (v *VPanel) Children() []Element {
	return v.childs
}

// AddChild adds the given children to the vpanel
func (v *VPanel) AddChild(e ...Element) {
	v.childs = append(v.childs, e...)
	for _, c := range e {
		c.SetParent(v)
	}
}

// Measure gets the "wanted" size of the element based on the available size
func (v *VPanel) Measure(availableWidth, availableHeight int) (width int, height int) {
	width, height = 0, 0
	v.heights = make(map[Element]int)
	v.widths = make(map[Element]int)
	for _, child := range v.childs {
		cw, ch := child.Measure(availableWidth, availableHeight-height)
		if cw > width {
			width = cw
		}
		height += ch
		v.heights[child] = ch
		v.widths[child] = cw
	}
	return
}

// Arrange sets the final size for the Element end tells it to Arrange itself
func (v *VPanel) Arrange(finalWidth, finalHeight int) {
	v.width, v.height = finalWidth, finalHeight

	for _, child := range v.childs {
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
	}
}

// Render renders the element on the given Renderer
func (v *VPanel) Render(rn Renderer) {
	yOff := 0
	for _, child := range v.childs {
		if yOff >= v.height {
			break
		}
		height := child.Height()
		rn.RenderChild(child, v.width, height, 0, yOff)
		yOff += height
	}
}

// Width returns the width of the vpanel
func (v *VPanel) Width() int {
	return v.width
}

// Height returns the height of the vpanel.
func (v *VPanel) Height() int {
	s := 0
	for _, child := range v.childs {
		s += child.Height()
	}
	return s
}

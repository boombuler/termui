package termui

type VPanel struct {
	BaseElement
	childs        []Element
	heights       map[Element]int
	widths        map[Element]int
	width, height int
}

var _ Element = new(VPanel) // Interface checking...

func NewVPanel() *VPanel {
	return new(VPanel)
}

func (v *VPanel) Name() string {
	return "VPanel"
}

func (v *VPanel) Children() []Element {
	return v.childs
}

func (v *VPanel) AddChild(e Element) {
	v.childs = append(v.childs, e)
	e.SetParent(v)
}

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

func (v *VPanel) Width() int {
	return v.width
}
func (v *VPanel) Height() int {
	s := 0
	for _, child := range v.childs {
		s += child.Height()
	}
	return s
}

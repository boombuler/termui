package termui

type Border struct {
	BaseElement
	child Element

	width, height int
}

var _ Element = new(Border) // Interface checking...

func (b *Border) Width() int {
	return b.width
}

func (b *Border) Height() int {
	return b.height
}

func (b *Border) Children() []Element {
	if b.child != nil {
		return []Element{b.child}
	}
	return []Element{}
}

func (b *Border) Name() string {
	return "border"
}

func (b *Border) Measure(availableWidth, availableHeight int) (width int, height int) {
	if b.child == nil {
		return 2, 2
	}

	cw, ch := b.child.Measure(availableWidth-2, availableHeight-2)
	grav := getGravity(b)
	if grav&horizontal == horizontal { // stretch
		width = availableWidth
	} else {
		width = cw + 2
	}
	if grav&vertical == vertical {
		height = availableHeight
	} else {
		height = ch + 2
	}
	return
}

func (b *Border) Arrange(finalWidth, finalHeight int) {
	if finalWidth > 2 && finalHeight > 2 && b.child != nil {
		b.child.Arrange(finalWidth-2, finalHeight-2)
	}
	b.width, b.height = finalWidth, finalHeight
}

func (b *Border) Render(r Renderer) {
	for x := 1; x < b.width-1; x++ {
		r.Set(x, 0, border_horizontal_line)
		r.Set(x, b.height-1, border_horizontal_line)

		for y := 1; y < b.height-1; y++ {
			r.Set(x, y, ' ')
		}
	}
	for y := 1; y < b.height-1; y++ {
		r.Set(0, y, border_vertical_line)
		r.Set(b.width-1, y, border_vertical_line)
	}

	r.Set(0, 0, border_top_left)
	r.Set(0, b.height-1, border_bottom_left)
	r.Set(b.width-1, 0, border_top_right)
	r.Set(b.width-1, b.height-1, border_bottom_right)

	if b.child != nil {
		r.RenderChild(b.child, b.width-2, b.height-2, 1, 1)
	}
}

func NewBorder(child Element) *Border {
	b := &Border{
		child: child,
	}
	if child != nil {
		child.SetParent(b)
	}
	return b
}

type TextBorder struct {
	*Border
	txt []rune
}

var _ Element = new(TextBorder) // Interface checking...

func NewTextBorder(txt string, child Element) *TextBorder {
	return &TextBorder{
		NewBorder(child),
		[]rune(txt),
	}
}

func (b *TextBorder) Render(rn Renderer) {
	b.Border.Render(rn)
	w := b.Width() - 2
	for i, r := range b.txt {
		if i > w {
			break
		}
		rn.Set(i+1, 0, r)
	}
}

func (b *TextBorder) SetText(txt string) {
	b.txt = []rune(txt)
}

func (b *TextBorder) Text() string {
	return string(b.txt)
}

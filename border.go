package termui

// Border draws a border around a given child element
type Border struct {
	BaseElement
	child Element

	width, height int
}

var _ Element = new(Border) // Interface checking...

// Width returns the width of the border
func (b *Border) Width() int {
	return b.width
}

// Height returns the height of the border
func (b *Border) Height() int {
	return b.height
}

// Children returns all nested child elements of the border
func (b *Border) Children() []Element {
	if b.child != nil {
		return []Element{b.child}
	}
	return []Element{}
}

// Name returns the constant name of the border struct for css styling
func (b *Border) Name() string {
	return "border"
}

// Measure gets the "wanted" size of the element based on the available size
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

// Arrange sets the final size for the Element end tells it to Arrange itself
func (b *Border) Arrange(finalWidth, finalHeight int) {
	if finalWidth > 2 && finalHeight > 2 && b.child != nil {
		b.child.Arrange(finalWidth-2, finalHeight-2)
	}
	b.width, b.height = finalWidth, finalHeight
}

// Render renders the element on the given Renderer
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

// NewBorder creates a new border element with the given child
func NewBorder(child Element) *Border {
	b := &Border{
		child: child,
	}
	if child != nil {
		child.SetParent(b)
	}
	return b
}

// TextBorder is a border with a text on the top left corner
type TextBorder struct {
	*Border
	txt []rune
}

var _ Element = new(TextBorder) // Interface checking...

// NewTextBorder creates a new TextBorder with a given text and child
func NewTextBorder(txt string, child Element) *TextBorder {
	return &TextBorder{
		NewBorder(child),
		[]rune(txt),
	}
}

// Render renders the element on the given Renderer
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

// SetText sets the text on the border
func (b *TextBorder) SetText(txt string) {
	b.txt = []rune(txt)
}

// Text returns the current text value
func (b *TextBorder) Text() string {
	return string(b.txt)
}

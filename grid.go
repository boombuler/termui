package termui

// Grid renders its children in a table like layout
type Grid struct {
	BaseElement

	// ColumnDefinitions contains the the column sizes for the grid
	ColumnDefinitions []int
	// RowDefinitions contains the the row sizes for the grid
	RowDefinitions []int

	rowHeights []int
	colWidths  []int
	children   []Element
	childPos   map[Element]GridPosition
}

var _ Element = new(Grid) // Interface checking...

const (
	// GridSizeAuto is a Grid Size that takes only that much space as needed
	GridSizeAuto int = 0
	// GridSizeAuto defines a Grid Size that takes as that much space as available after layout all others.
	// The total width is distributed between all rows / columns with Star definitions
	// For example you can have one column with 1*GridSizeStar and one with 2*GridSizeStar. In this case
	// the second column will be twice as wide as the first one.
	GridSizeStar int = -1
)

// GridPosition defines the location of an element within a grid
type GridPosition struct {
	// Column contains the column of the element
	Column int
	// Row contains the row of element
	Row int
	// ColumnSpan of the element. Auto-Width columns doesn't take spanned elements into account
	ColumnSpan int
	// RowSpan of the element. Auto-Height rows doesn't take spanned elements into account
	RowSpan int
}

// NewGrid returns a new Grid.
func NewGrid(rowDefinitions, colDefinitions []int) *Grid {
	return &Grid{
		ColumnDefinitions: colDefinitions,
		RowDefinitions:    rowDefinitions,
		childPos:          make(map[Element]GridPosition),
	}
}

func (g *Grid) AddChild(e Element, pos GridPosition) {
	g.children = append(g.children, e)
	if pos.Column < 0 || pos.Column >= len(g.ColumnDefinitions) {
		pos.Column = 0
	}
	if pos.Row < 0 || pos.Row >= len(g.RowDefinitions) {
		pos.Row = 0
	}
	if pos.RowSpan < 1 || pos.RowSpan > len(g.RowDefinitions) {
		pos.RowSpan = 1
	}
	if pos.ColumnSpan < 1 || pos.ColumnSpan > len(g.ColumnDefinitions) {
		pos.ColumnSpan = 1
	}
	g.childPos[e] = pos
}

// Name returns the constant name of the grid for css styling.
func (g *Grid) Name() string {
	return "grid"
}

// Children returns all nested elements.
func (g *Grid) Children() []Element {
	return g.children
}

func sumIntSlice(v []int) int {
	s := 0
	for _, val := range v {
		s += val
	}
	return s
}

// Height returns the height of the element
func (g *Grid) Height() int {
	return sumIntSlice(g.rowHeights)
}

// Width returns the width of the border
func (g *Grid) Width() int {
	return sumIntSlice(g.colWidths)
}

func (g *Grid) measureColumnAuto(col int) int {
	width := 0

	for c, p := range g.childPos {
		if p.Column == col && p.ColumnSpan == 1 {
			w, _ := c.Measure(0, 0)
			if w > width {
				width = w
			}
		}
	}
	return width
}

func (g *Grid) measureRowAuto(row int) int {
	height := 0

	for c, p := range g.childPos {
		if p.Row == row && p.RowSpan == 1 {
			_, h := c.Measure(0, 0)
			if h > height {
				height = h
			}
		}
	}
	return height
}

func (g *Grid) measureGridLengths(availableWidth, availableHeight int) ([]int, []int) {
	reqWidth := 0
	widths := make([]int, len(g.ColumnDefinitions))
	for col, colW := range g.ColumnDefinitions {
		if colW > GridSizeAuto {
			reqWidth += colW
			widths[col] = colW
		} else if colW == GridSizeAuto {
			widths[col] = g.measureColumnAuto(col)
			reqWidth += widths[col]
		}
	}

	reqHeight := 0
	heights := make([]int, len(g.RowDefinitions))
	for row, rowH := range g.RowDefinitions {
		if rowH > GridSizeAuto {
			reqHeight += rowH
			heights[row] = rowH
		} else if rowH == GridSizeAuto {
			heights[row] = g.measureRowAuto(row)
			reqHeight += heights[row]
		}
	}

	starWTotal := availableWidth - reqWidth
	starWTotalCnt := 0
	for _, colW := range g.ColumnDefinitions {
		if colW < GridSizeAuto {
			starWTotalCnt += (colW / GridSizeStar)
		}
	}

	if starWTotal > 0 && starWTotalCnt > 0 {
		for col, colW := range g.ColumnDefinitions {
			if colW < GridSizeAuto {
				perc := float64(colW/GridSizeStar) / float64(starWTotalCnt)
				widths[col] = int(float64(starWTotal) * perc)
			}
		}
	}

	starHTotal := availableHeight - reqHeight
	starHTotalCnt := 0
	for _, rowH := range g.RowDefinitions {
		if rowH < GridSizeAuto {
			starHTotalCnt += (rowH / GridSizeStar)
		}
	}
	if starHTotal > 0 && starHTotalCnt > 0 {
		for row, rowH := range g.RowDefinitions {
			if rowH < GridSizeAuto {
				perc := float64(rowH/GridSizeStar) / float64(starHTotalCnt)
				heights[row] = int(float64(starHTotal) * perc)
			}
		}
	}
	return widths, heights
}

// Measure gets the "wanted" size of the element based on the available size
func (g *Grid) Measure(availableWidth, availableHeight int) (width int, height int) {
	widths, heights := g.measureGridLengths(availableWidth, availableHeight)
	for child, pos := range g.childPos {
		child.Measure(
			sumIntSlice(widths[pos.Column:pos.Column+pos.ColumnSpan]),
			sumIntSlice(heights[pos.Row:pos.Row+pos.RowSpan]))
	}
	return sumIntSlice(widths), sumIntSlice(heights)
}

// Arrange sets the final size for the Element end tells it to Arrange itself
func (g *Grid) Arrange(finalWidth, finalHeight int) {
	g.colWidths, g.rowHeights = g.measureGridLengths(finalWidth, finalHeight)
	for c, p := range g.childPos {
		c.Arrange(sumIntSlice(g.colWidths[p.Column:p.Column+p.ColumnSpan]),
			sumIntSlice(g.rowHeights[p.Row:p.Row+p.RowSpan]))
	}
}

// Render renders the element on the given Renderer
func (g *Grid) Render(r Renderer) {
	for c, p := range g.childPos {
		r.RenderChild(c,
			sumIntSlice(g.colWidths[p.Column:p.Column+p.ColumnSpan]),
			sumIntSlice(g.rowHeights[p.Row:p.Row+p.RowSpan]),
			sumIntSlice(g.colWidths[:p.Column]),
			sumIntSlice(g.rowHeights[:p.Row]))
	}
}
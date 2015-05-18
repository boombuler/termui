package termui

import (
	"testing"
)

type fixedSizeBlock struct {
	BaseElement
	width, height        int
	ArrangedW, ArrangedH int
}

var _ Element = new(fixedSizeBlock)

func newFixedBlock(w, h int) *fixedSizeBlock {
	return &fixedSizeBlock{
		width:  w,
		height: h,
	}
}
func (f *fixedSizeBlock) Name() string {
	return ""
}
func (f *fixedSizeBlock) Render(r Renderer) {
}
func (f *fixedSizeBlock) Arrange(w, h int) {
	f.ArrangedW, f.ArrangedH = w, h
}
func (f *fixedSizeBlock) Measure(w, h int) (int, int) {
	return f.width, f.height
}
func (f *fixedSizeBlock) Children() []Element {
	return []Element{}
}

func (f *fixedSizeBlock) Width() int {
	return f.width
}

func (f *fixedSizeBlock) Height() int {
	return f.height
}

func Test_GridLayoutColumnsAuto(t *testing.T) {
	g := NewGrid([]int{GridSizeAuto}, []int{GridSizeAuto, 10})
	g.AddChild(newFixedBlock(100, 10), GridPosition{0, 0, 1, 1})
	w, h := g.Measure(0, 0)
	if w != 110 {
		t.Errorf("Grid Layout Width failed (got %v expected %v)", w, 110)
	}
	if h != 10 {
		t.Errorf("Grid Layout Height failed (got %v expected %v)", h, 10)
	}
	g.AddChild(newFixedBlock(150, 20), GridPosition{0, 0, 2, 1})
	w, h = g.Measure(0, 0)
	if w != 150 {
		t.Errorf("Grid Layout Width failed (got %v expected %v)", w, 150)
	}
	if h != 20 {
		t.Errorf("Grid Layout Height failed (got %v expected %v)", h, 20)
	}
}

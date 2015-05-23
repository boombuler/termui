package termui

import (
	"fmt"
	"github.com/boombuler/termui/css"
	"testing"
)

type fixedSizeBlock struct {
	BaseElement
	width, height        int
	ArrangedW, ArrangedH int
	MeasuredW, MeasuredH int
}

func (f fixedSizeBlock) String() string {
	return fmt.Sprintf("Fixed { W: %d, H: %d, AW: %d, AH: %d, MW: %d, MH: %d }", f.width, f.height, f.ArrangedW, f.ArrangedH, f.MeasuredW, f.MeasuredH)
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
	f.MeasuredW, f.MeasuredH = w, h
	return f.width, f.height
}
func (f *fixedSizeBlock) Children() []css.Styleable {
	return nil
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

func Test_GridChildWithMargin(t *testing.T) {
	g := NewGrid([]int{GridSizeAuto}, []int{GridSizeAuto})
	child := newFixedBlock(10, 10)
	g.AddChild(child, GridPosition{0, 0, 1, 1})
	w, h := g.Measure(0, 0)
	if w != 10 || h != 10 {
		t.Errorf("Invalid Gridsize: %vx%v expected %vx%v child: %v", w, h, 10, 10, child)
	}

	child.SetProperty(MarginProperty.Property, Thickness{5, 5, 5, 5})
	css.ClearCache()

	w, h = g.Measure(0, 0)
	if w != 20 || h != 20 {
		t.Errorf("Invalid Gridsize: %vx%v expected %vx%v child: %v", w, h, 20, 20, child)
	}
}

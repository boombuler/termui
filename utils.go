package termui

// MeasureChild measures the child and takes margin in account
func MeasureChild(e Element, availableWidth, availableHeight int) (int, int) {
	if availableHeight < 0 {
		availableHeight = 0
	}
	if availableWidth < 0 {
		availableWidth = 0
	}
	m := MarginProperty.Get(e)
	mw := m.Left + m.Right
	mh := m.Top + m.Bottom
	if availableWidth > mw {
		availableWidth -= mw
	}
	if availableHeight > mh {
		availableHeight -= mh
	}
	w, h := e.Measure(availableWidth, availableHeight)
	return w + mw, h + mh
}

// ArrangeChild arranges the child and takes margin in account
func ArrangeChild(e Element, finalWidth, finalHeight int) {
	m := MarginProperty.Get(e)
	mw := m.Left + m.Right
	mh := m.Top + m.Bottom
	e.Arrange(finalWidth-mw, finalHeight-mh)
}

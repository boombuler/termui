package termui

// MeasureChild measures the child and takes margin in account
func MeasureChild(e Element, availableWidth, availableHeight int) (int, int) {
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

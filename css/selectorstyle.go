package css

// SelectorStyle is a combination of a selector and a style.
type SelectorStyle struct {
	Selector
	Style Style
}

// SelectorStyles is a slice of SelectorStyle which can be sorted.
type SelectorStyles []SelectorStyle

// Less tests if the Selector on position i is less important then the on on position j.
func (s SelectorStyles) Less(i, j int) bool {
	si, sj := s[i].weight(), s[j].weight()
	if si.A < sj.A {
		return true
	} else if si.A == sj.A {
		if si.B < sj.B {
			return true
		} else if si.B == sj.B {
			if si.C < sj.C {
				return true
			} else if si.C == sj.C && si.D < sj.D {
				return true
			}
		}
	}
	return false
}

// Len returns the length of the slice.
func (s SelectorStyles) Len() int {
	return len(s)
}

// Swap exchanges the elements at position i and j
func (s SelectorStyles) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

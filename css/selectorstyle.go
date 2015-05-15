package css

type SelectorStyle struct {
	Selector
	Style Style
}

type SelectorStyles []SelectorStyle

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

func (s SelectorStyles) Len() int {
	return len(s)
}

func (s SelectorStyles) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

package css

type Selector interface {
	Matches(Styleable) bool
	weight() selectorWeight
}

type selectorWeight struct {
	A, B, C, D int
}

func (sw selectorWeight) Add(other selectorWeight) selectorWeight {
	return selectorWeight{
		sw.A + other.A,
		sw.B + other.B,
		sw.C + other.C,
		sw.D + other.D,
	}
}

var (
	AnySelector  Selector = anySelector{}
	BodySelector Selector = bodySelector{}
)

type bodySelector struct{}

func (es bodySelector) Matches(s Styleable) bool {
	return s == nil
}

func (es bodySelector) weight() selectorWeight {
	return selectorWeight{0, 0, 0, 0}
}

type anySelector struct{}

func (as anySelector) Matches(s Styleable) bool {
	return true
}
func (as anySelector) weight() selectorWeight {
	return selectorWeight{0, 0, 0, 0}
}

type NameSelector string

func (ns NameSelector) Matches(s Styleable) bool {
	return s != nil && s.Name() == string(ns)
}
func (ns NameSelector) weight() selectorWeight {
	return selectorWeight{0, 0, 0, 1}
}

type PseudoClassSelector string

func (pcs PseudoClassSelector) Matches(s Styleable) bool {
	if s == nil {
		return false
	}
	for _, class := range s.PseudoClasses() {
		if class == string(pcs) {
			return true
		}
	}
	return false
}

func (pcs PseudoClassSelector) weight() selectorWeight {
	return selectorWeight{0, 0, 1, 0}
}

type ClassSelector string

func (cs ClassSelector) Matches(s Styleable) bool {
	if s == nil {
		return false
	}
	for _, class := range s.Classes() {
		if class == string(cs) {
			return true
		}
	}
	return false
}

func (cs ClassSelector) weight() selectorWeight {
	return selectorWeight{0, 0, 1, 0}
}

type IdSelector string

func (is IdSelector) Matches(s Styleable) bool {
	return s != nil && s.Id() == string(is)
}

func (is IdSelector) weight() selectorWeight {
	return selectorWeight{0, 1, 0, 0}
}

type ElementSelector struct {
	elem Styleable
}

func NewElementSelector(elem Styleable) ElementSelector {
	return ElementSelector{elem}
}

func (es ElementSelector) Matches(s Styleable) bool {
	return es.elem == s
}

func (es ElementSelector) weight() selectorWeight {
	return selectorWeight{1, 0, 0, 0}
}

type ParentSelector struct {
	Parent, Child Selector
}

func (ps ParentSelector) Matches(s Styleable) bool {
	if ps.Child.Matches(s) {
		if s != nil {
			s = s.Parent()
			return ps.Parent.Matches(s)
		}
	}
	return false
}

func (ps ParentSelector) weight() selectorWeight {
	return ps.Child.weight().Add(ps.Parent.weight())
}

type AndSelector []Selector

func (as AndSelector) Matches(s Styleable) bool {
	for _, sel := range as {
		if !sel.Matches(s) {
			return false
		}
	}
	return true
}

func (as AndSelector) weight() selectorWeight {
	var sum selectorWeight

	for _, sel := range as {
		sum = sum.Add(sel.weight())
	}

	return sum
}

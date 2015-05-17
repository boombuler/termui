package css

// Selector defines a CSS selector which can be matched against styleable elements.
type Selector interface {
	// Matches returns true if the selector matches the styleable element.
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
	// AnySelector is a selector which matches any element.
	AnySelector Selector = anySelector{}
	// BodySelector is a selector which matches a virtual "body" element.
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

// NameSelector is a selector which matches the name of the styleable element.
type NameSelector string

// Matches returns true if the selector matches the styleable element.
func (ns NameSelector) Matches(s Styleable) bool {
	return s != nil && s.Name() == string(ns)
}
func (ns NameSelector) weight() selectorWeight {
	return selectorWeight{0, 0, 0, 1}
}

// ClassSelector matches an element if it has the given class.
type ClassSelector string

// Matches returns true if the selector matches the styleable element.
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

// IDSelector matches an element by its id.
type IDSelector string

// Matches returns true if the selector matches the styleable element.
func (is IDSelector) Matches(s Styleable) bool {
	return s != nil && s.ID() == string(is)
}

func (is IDSelector) weight() selectorWeight {
	return selectorWeight{0, 1, 0, 0}
}

type elementSelector struct {
	elem Styleable
}

// Matches returns true if the selector matches the styleable element.
func (es elementSelector) Matches(s Styleable) bool {
	return es.elem == s
}

func (es elementSelector) weight() selectorWeight {
	return selectorWeight{1, 0, 0, 0}
}

// ParentSelector matches if the Parent field matches the direkt parent of the element and the Child field matches the element.
type ParentSelector struct {
	// Parent is tested against the direct parent of the element.
	Parent,
	// Child is tested against the element.
	Child Selector
}

// Matches returns true if the selector matches the styleable element.
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

// AnyParentSelector is a selector which matches if any parent matches the parent selector.
type AnyParentSelector struct {
	// Parent selector which should match a parent at any level.
	Parent,
	// Child selector which must be matched by the element.
	Child Selector
}

// Matches returns true if the selector matches the styleable element.
func (ps AnyParentSelector) Matches(s Styleable) bool {
	if ps.Child.Matches(s) {
		for s != nil {
			s = s.Parent()
			if ps.Parent.Matches(s) {
				return true
			}
		}
	}
	return false
}

func (ps AnyParentSelector) weight() selectorWeight {
	return ps.Child.weight().Add(ps.Parent.weight())
}

// AndSelector is a selector which matches if all nested selectors match.
type AndSelector []Selector

// Matches returns true if the selector matches the styleable element.
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

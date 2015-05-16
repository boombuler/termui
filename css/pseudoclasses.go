package css

// PseudoClassSelector is a CSS-Selector for Pseudoclass matching.
type PseudoClassSelector string

// PseudoClassMatcher is used to tell the styleengine if a styleable has a specific pseudo class.
type PseudoClassMatcher interface {
	// Matches checks if the styleable have the pseudoclass
	Matches(s Styleable) bool
}

// PseudoClassMatcherFunc implements the PseudoClassMatcher interface
type PseudoClassMatcherFunc func(s Styleable) bool

// Matches checks if the styleable have the pseudoclass
func (mf PseudoClassMatcherFunc) Matches(s Styleable) bool {
	return mf(s)
}

var pseudoClassMatchers = make(map[string]PseudoClassMatcher)

// SetPseudoClassMatcher sets the PseudoClassMatcher for a pseudoclass with a given name.
func SetPseudoClassMatcher(class string, matcher PseudoClassMatcher) {
	if matcher == nil {
		delete(pseudoClassMatchers, class)
	} else {
		pseudoClassMatchers[class] = matcher
	}
}

// Matches checks if the styleable have the pseudoclass
func (pcs PseudoClassSelector) Matches(s Styleable) bool {
	if m, ok := pseudoClassMatchers[string(pcs)]; ok {
		return m.Matches(s)
	}
	return false
}

func (pcs PseudoClassSelector) weight() selectorWeight {
	return selectorWeight{0, 0, 1, 0}
}

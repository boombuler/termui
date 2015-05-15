package css

type PseudoClassSelector string

type PseudoClassMatcher interface {
	Matches(s Styleable) bool
}

type PseudoClassMatcherFunc func(s Styleable) bool

func (mf PseudoClassMatcherFunc) Matches(s Styleable) bool {
	return mf(s)
}

var pseudoClassMatchers map[string]PseudoClassMatcher = make(map[string]PseudoClassMatcher)

func SetPseudoClassMatcher(class string, matcher PseudoClassMatcher) {
	if matcher == nil {
		delete(pseudoClassMatchers, class)
	} else {
		pseudoClassMatchers[class] = matcher
	}
}

func (pcs PseudoClassSelector) Matches(s Styleable) bool {
	if m, ok := pseudoClassMatchers[string(pcs)]; ok {
		return m.Matches(s)
	}
	return false
}

func (pcs PseudoClassSelector) weight() selectorWeight {
	return selectorWeight{0, 0, 1, 0}
}

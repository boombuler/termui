package css

import (
	"sort"
)

var (
	userAgentStyles SelectorStyles
	designerStyle   SelectorStyles
	userStyle       SelectorStyles
)

// AddUserAgentStyle sets the default style for an UI element
func AddUserAgentStyle(sel Selector, s Style) {
	userAgentStyles = append(userAgentStyles, SelectorStyle{sel, s})
}

func applyMatchingStyles(dict []SelectorStyle, s Styleable, curStyle Style) Style {
	var matching SelectorStyles
	for _, sSt := range dict {
		if sSt.Matches(s) {
			matching = append(matching, sSt)
		}
	}
	sort.Sort(matching)
	for _, sSt := range matching {
		curStyle = curStyle.Merge(sSt.Style)
	}
	return curStyle
}

// Get returns the calculated style for the given Styleable element.
func Get(s Styleable) Style {
	var baseStyle Style
	if s != nil {
		baseStyle = Get(s.Parent()).Inherit()
	}

	return applyMatchingStyles(userStyle, s,
		applyMatchingStyles(designerStyle, s,
			applyMatchingStyles(userAgentStyles, s, baseStyle)))
}

// SetUserStyles sets the style overrides from the user. This replaces all other user styles which were added before.
func SetUserStyles(styles SelectorStyles) {
	userStyle = styles
}

// SetDesignerStyles sets the designer overrides for the program. This replaces all other designer styles which were added before.
func SetDesignerStyles(styles SelectorStyles) {
	designerStyle = styles
}

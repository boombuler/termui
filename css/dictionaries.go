package css

import (
	"sort"
)

var (
	userAgentStyles SelectorStyles
	designerStyle   SelectorStyles
	userStyle       SelectorStyles
)

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

func Get(s Styleable) Style {
	var baseStyle Style
	if s != nil {
		baseStyle = Get(s.Parent()).Inherit()
	}

	return applyMatchingStyles(userStyle, s,
		applyMatchingStyles(designerStyle, s,
			applyMatchingStyles(userAgentStyles, s, baseStyle)))
}

func SetUserStyles(styles SelectorStyles) {
	userStyle = styles
}

func SetDesignerStyles(styles SelectorStyles) {
	designerStyle = styles
}

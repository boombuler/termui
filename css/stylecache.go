package css

import (
	"sort"
)

var cache = make(map[Styleable]Style)

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

func buildStyle(s Styleable) Style {
	var baseStyle Style
	var elemStyleDict []SelectorStyle
	if s != nil {
		baseStyle = Get(s.Parent()).Inherit()
		elemStyle := s.ElementStyle()
		if elemStyle != nil {
			elemStyleDict = []SelectorStyle{
				SelectorStyle{elementSelector{s}, elemStyle},
			}
		}
	}

	return applyMatchingStyles(elemStyleDict, s,
		applyMatchingStyles(userStyle, s,
			applyMatchingStyles(designerStyle, s,
				applyMatchingStyles(userAgentStyles, s, baseStyle))))
}

func ClearCache() {
	cache = make(map[Styleable]Style)
}

// Get returns the calculated style for the given Styleable element.
func Get(s Styleable) Style {
	style, ok := cache[s]
	if ok {
		return style
	} else {
		cache[s] = buildStyle(s)
		return cache[s]
	}
}

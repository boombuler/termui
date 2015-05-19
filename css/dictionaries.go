package css

var (
	userAgentStyles SelectorStyles
	designerStyle   SelectorStyles
	userStyle       SelectorStyles
)

// AddUserAgentStyle sets the default style for an UI element
func AddUserAgentStyle(sel Selector, s Style) {
	userAgentStyles = append(userAgentStyles, SelectorStyle{sel, s})
	ClearCache()
}

// SetUserStyles sets the style overrides from the user. This replaces all other user styles which were added before.
func SetUserStyles(styles SelectorStyles) {
	userStyle = styles
	ClearCache()
}

// SetDesignerStyles sets the designer overrides for the program. This replaces all other designer styles which were added before.
func SetDesignerStyles(styles SelectorStyles) {
	designerStyle = styles
	ClearCache()
}

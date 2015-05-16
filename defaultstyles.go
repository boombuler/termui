package termui

import (
	"github.com/boombuler/termui/css"
	"github.com/nsf/termbox-go"
)

func init() {
	css.AddUserAgentStyle(css.BodySelector, css.Style{
		BackgroundProperty.Property: termbox.ColorDefault,
		ForegroundProperty.Property: termbox.ColorDefault,
	})

	textbox := css.NameSelector("textbox")
	css.AddUserAgentStyle(textbox, css.Style{
		BackgroundProperty.Property: termbox.ColorWhite,
		ForegroundProperty.Property: termbox.ColorBlack,
	})

	focused := css.PseudoClassSelector(pclassFocused)
	css.AddUserAgentStyle(css.AndSelector{textbox, focused}, css.Style{
		BackgroundProperty.Property: termbox.ColorBlue,
		ForegroundProperty.Property: termbox.ColorWhite,
	})
}

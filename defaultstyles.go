package termui

import (
	"github.com/boombuler/termui/css"
	"github.com/nsf/termbox-go"
)

func init() {
	css.AddUserAgentStyle(css.BodySelector, css.Style{
		BackgroundProperty: termbox.ColorDefault,
		ForegroundProperty: termbox.ColorDefault,
	})

	textbox := css.NameSelector("textbox")
	css.AddUserAgentStyle(textbox, css.Style{
		BackgroundProperty: termbox.ColorWhite,
		ForegroundProperty: termbox.ColorBlack,
	})

	focused := css.PseudoClassSelector(pclass_Focused)
	css.AddUserAgentStyle(css.AndSelector{textbox, focused}, css.Style{
		BackgroundProperty: termbox.ColorBlue,
		ForegroundProperty: termbox.ColorWhite,
	})
}

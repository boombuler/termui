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
}

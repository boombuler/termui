package main

//go:generate termcssgen -i=termtest.css -o=style_gen.go -p=main -m=defaultStyle

import (
	"github.com/boombuler/termui"
	"github.com/boombuler/termui/css"
	"github.com/nsf/termbox-go"
)

func main() {
	css.SetDesignerStyles(defaultStyle())

	tx := termui.NewText("Hallo Welt")
	innerBox := termui.NewBorder(tx)
	box := termui.NewTextBorder("Foobar", innerBox)

	termui.Start(box)
	go func() {
		for ev := range termui.Events {
			if ev.Type == termbox.EventKey {
				if ev.Key == termbox.KeyEsc {
					termui.Stop()
				} else {
					tx.SetText(string(append([]rune(tx.Text()), ev.Ch)))
					termui.Update()
				}
			}
		}
	}()

	termui.Wait()
}

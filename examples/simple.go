package main

//go:generate termcssgen -i=simple.css -o=style_gen.go -p=main -m=defaultStyle

import (
	"github.com/boombuler/termui"
	"github.com/boombuler/termui/css"
	"github.com/nsf/termbox-go"
)

func main() {
	css.SetDesignerStyles(defaultStyle())

	tx := termui.NewText("Hallo Welt")

	vInner := termui.NewVPanel()
	vInner.AddChild(tx)
	vInner.AddChild(termui.NewTextBox())
	vInner.AddChild(termui.NewTextBox())

	innerBox := termui.NewBorder(vInner)

	box := termui.NewTextBorder("Foobar", innerBox)

	termui.Start(box)
	go func() {
		for ev := range termui.Events {
			if ev.Type == termbox.EventKey {
				if ev.Key == termbox.KeyEsc {
					termui.Stop()
				}
			}
		}
	}()

	termui.Wait()
}

package main

import (
	"github.com/boombuler/termui"
	"github.com/nsf/termbox-go"
)

func main() {
	grd := termui.NewGrid([]int{5, 1, termui.GridSizeAuto},
		[]int{
			termui.GridSizeStar,
			2,
			2 * termui.GridSizeStar,
			2,
			termui.GridSizeAuto,
		})

	grd.AddChild(termui.NewBorder(termui.NewText("One Star")), 0, 0)
	grd.AddChild(termui.NewBorder(termui.NewText("Two Star")), 0, 2)
	grd.AddChild(termui.NewBorder(termui.NewText("Auto")), 0, 4)
	grd.AddChild(termui.NewBorder(termui.NewText("Press Esc to exit!")), 2, 2)

	termui.Start(grd)
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

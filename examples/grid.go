package main

import (
	"github.com/boombuler/termui"
	"github.com/nsf/termbox-go"
)

func main() {
	grd := termui.NewGrid(
		[]int{
			termui.GridSizeStar,
			2,
			termui.GridSizeAuto,
			2,
			2 * termui.GridSizeStar,
		}, []int{5, 1, termui.GridSizeAuto})

	grd.AddChild(termui.NewBorder(termui.NewText("One Star")), termui.GridPosition{0, 0, 1, 1})
	grd.AddChild(termui.NewBorder(termui.NewText("Auto")), termui.GridPosition{2, 0, 1, 1})
	grd.AddChild(termui.NewBorder(termui.NewText("Two Star")), termui.GridPosition{4, 0, 1, 1})
	grd.AddChild(termui.NewBorder(termui.NewText("Press Esc to exit!")), termui.GridPosition{0, 2, 5, 1})

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

package termui

import (
	"github.com/nsf/termbox-go"
)

type intputManager struct {
	root    Element
	current FocusElement
}

func newInputManager(root Element) *intputManager {
	im := &intputManager{root: root}
	im.Next()
	return im
}

func (im *intputManager) Next() {
	itmStack := []Element{im.root}
	items := make([]FocusElement, 0)
	for len(itmStack) > 0 {
		lastIdx := len(itmStack) - 1
		cur := itmStack[lastIdx]
		if fcur, ok := cur.(FocusElement); ok {
			items = append(items, fcur)
		}
		itmStack = append(itmStack[:lastIdx], cur.Children()...)
	}
	if im.current != nil {
		im.current.SetFocused(false)
	}
	if len(items) > 0 {
		curIdx := -1
		for i, c := range items {
			if c == im.current {
				curIdx = i
				break
			}
		}
		curIdx = (curIdx + 1) % len(items)
		im.current = items[curIdx]
		if im.current != nil {
			im.current.SetFocused(true)
		}
	} else {
		im.current = nil
	}
}

func (im *intputManager) DispatchEvent(e termbox.Event) bool {
	if e.Type == termbox.EventKey {
		if im.current != nil {
			if im.current.HandleKey(e.Key, e.Ch) {
				return true
			}
		}
		if e.Key == termbox.KeyTab {
			im.Next()
			return im.current != nil
		}
	}
	return false
}

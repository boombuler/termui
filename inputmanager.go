package termui

import (
	"github.com/boombuler/termui/css"
	"github.com/nsf/termbox-go"
)

const pclassFocused = "focused"

type intputManager struct {
	root    Element
	current FocusElement
}

func newInputManager(root Element) *intputManager {
	im := &intputManager{root: root}
	im.Next()

	css.SetPseudoClassMatcher(pclassFocused, css.PseudoClassMatcherFunc(func(s css.Styleable) bool {
		return im.current == s
	}))
	return im
}

func (im *intputManager) TrySetFocusTo(fe FocusElement) {
	if fe == nil {
		return
	}
	for _, e := range im.getFocusableItems() {
		if e == fe {
			im.current.SetFocused(false)
			im.current = fe
			im.current.SetFocused(true)
			return
		}
	}
}

func (im *intputManager) getFocusableItems() []FocusElement {
	itmStack := []Element{im.root}
	var items []FocusElement
	for len(itmStack) > 0 {
		lastIdx := len(itmStack) - 1
		cur := itmStack[lastIdx]
		childs := cur.Children()
		revChilds := make([]Element, 0)
		for _, c := range childs {
			el, ok := c.(Element)
			if ok {
				revChilds = append([]Element{el}, revChilds...)
			}
		}

		if fcur, ok := cur.(FocusElement); ok {
			items = append(items, fcur)
		}
		itmStack = append(itmStack[:lastIdx], revChilds...)
	}
	return items
}
func (im *intputManager) Next() {
	items := im.getFocusableItems()
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
		var kh KeyHandler = im.current
		for kh != nil {
			if kh.HandleKey(e.Key, e.Ch) {
				return true
			}
			p, ok := kh.Parent().(Element)
			kh = nil
			for p != nil && ok {
				if k, ok := p.(KeyHandler); ok {
					kh = k
					break
				}
				p, ok = p.Parent().(Element)
			}
		}
		if e.Key == termbox.KeyTab {
			im.Next()
			return im.current != nil
		}
	}
	return false
}

package termui

type ElementCollection struct {
	items []Element
}

func (c *ElementCollection) Add(e ...Element) {
	c.items = append(c.items, e...)
}

func (c *ElementCollection) IndexOf(e Element) int {
	for i, itm := range c.items {
		if itm == e {
			return i
		}
	}
	return -1
}

func (c *ElementCollection) Remove(e Element) {
	c.RemoveAt(c.IndexOf(e))
}

func (c *ElementCollection) RemoveAt(i int) {
	if i >= 0 && i < len(c.items) {
		c.items = append(c.items[:i], c.items[i+1:]...)
	}
}

func (c *ElementCollection) Clear() {
	c.items = nil
}

func (c *ElementCollection) Len() int {
	return len(c.items)
}

func (c *ElementCollection) At(i int) Element {
	return c.items[i]
}

func (c *ElementCollection) Items() []Element {
	return c.items
}

package termui

import (
	"testing"
)

func Test_Next(t *testing.T) {
	vPanel := NewVPanel()
	tb1 := NewTextBox()
	tb2 := NewTextBox()
	tb3 := NewTextBox()
	vPanel.AddChild(tb1)
	vPanel.AddChild(tb2)
	vPanel.AddChild(tb3)

	getTBNum := func(t Element) int {
		if t == tb1 {
			return 1
		}
		if t == tb2 {
			return 2
		}
		return 3
	}
	im := newInputManager(vPanel)

	if im.current != tb1 {
		t.Errorf("Test 1: TB %v should be focused but TB %v is focused!", getTBNum(tb1), getTBNum(im.current))
	}
	im.Next()
	if im.current != tb2 {
		t.Errorf("Test 2: TB %v should be focused but TB %v is focused!", getTBNum(tb2), getTBNum(im.current))
	}
	im.Next()
	if im.current != tb3 {
		t.Errorf("Test 3: TB %v should be focused but TB %v is focused!", getTBNum(tb3), getTBNum(im.current))
	}
	im.Next()
	if im.current != tb1 {
		t.Errorf("Test 4: TB %v should be focused but TB %v is focused!", getTBNum(tb1), getTBNum(im.current))
	}
}

package termui

import (
	"fmt"
	"github.com/boombuler/termui/css"
)

var MarginProperty = NewThicknessProperty("margin")

type Thickness struct {
	Top, Right, Bottom, Left int
}

type ThicknessPropertyDefinition struct {
	Property                 *css.Property
	top, right, bottom, left *css.Property
}

func (tp *ThicknessPropertyDefinition) Get(e Element) (res Thickness) {
	style := css.Get(e)
	if all, ok := style[tp.Property]; ok {
		res = all.(Thickness)
	}

	if t, ok := style[tp.top]; ok {
		res.Top = t.(int)
	}
	if r, ok := style[tp.right]; ok {
		res.Right = r.(int)
	}
	if b, ok := style[tp.bottom]; ok {
		res.Bottom = b.(int)
	}
	if l, ok := style[tp.left]; ok {
		res.Left = l.(int)
	}
	return
}

func NewThicknessProperty(name string) *ThicknessPropertyDefinition {
	return &ThicknessPropertyDefinition{
		Property: css.NewProperty(name, css.PropertyBehaviorDefault, thicknessConverter{}),
		top:      css.NewProperty(fmt.Sprintf("%v-top", name), css.PropertyBehaviorDefault, singleIntConv{}),
		right:    css.NewProperty(fmt.Sprintf("%v-right", name), css.PropertyBehaviorDefault, singleIntConv{}),
		bottom:   css.NewProperty(fmt.Sprintf("%v-bottom", name), css.PropertyBehaviorDefault, singleIntConv{}),
		left:     css.NewProperty(fmt.Sprintf("%v-left", name), css.PropertyBehaviorDefault, singleIntConv{}),
	}
}

type thicknessConverter struct{}

func (tc thicknessConverter) FromInt(i int) (interface{}, error) {
	return Thickness{i, i, i, i}, nil
}

func (tc thicknessConverter) FromStrings(s ...string) (interface{}, error) {
	intValues := make([]int, len(s))
	for i, sv := range s {
		var iv int
		if _, err := fmt.Sscanf(sv, "%d", &iv); err != nil {
			return nil, err
		} else {
			intValues[i] = iv
		}
	}
	switch len(intValues) {
	case 1:
		return Thickness{intValues[0], intValues[0], intValues[0], intValues[0]}, nil
	case 2:
		return Thickness{intValues[0], intValues[1], intValues[0], intValues[1]}, nil
	case 3:
		return Thickness{intValues[0], intValues[1], intValues[2], intValues[1]}, nil
	case 4:
		return Thickness{intValues[0], intValues[1], intValues[2], intValues[3]}, nil
	default:
		return nil, fmt.Errorf("Invalid argument count.")
	}
}

type singleIntConv struct{}

func (sic singleIntConv) FromInt(i int) (interface{}, error) {
	return i, nil
}

func (sic singleIntConv) FromStrings(s ...string) (interface{}, error) {
	if len(s) != 1 {
		return nil, fmt.Errorf("Unexpected values: %v expected number", s)
	}
	var i int
	if _, err := fmt.Sscanf(s[0], "%d", &i); err != nil {
		return nil, err
	} else {
		return i, nil
	}
}

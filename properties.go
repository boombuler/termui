package termui

import (
	"fmt"
	"github.com/boombuler/termui/css"
	"github.com/nsf/termbox-go"
)

var (
	BackgroundProperty = css.NewProperty("background", css.PropertyBehaviorInherit, attributeConverter{})
	ForegroundProperty = css.NewProperty("color", css.PropertyBehaviorInherit, attributeConverter{})
	GravityProperty    = css.NewProperty("gravity", css.PropertyBehaviorDefault, gravityConverter{})
)

type attributeConverter struct{}

func (ac attributeConverter) FromInt(i int) (interface{}, error) {
	return termbox.Attribute(i), nil
}

func (ac attributeConverter) FromStrings(s ...string) (interface{}, error) {
	values := map[string]termbox.Attribute{
		"default":   termbox.ColorDefault,
		"black":     termbox.ColorBlack,
		"red":       termbox.ColorRed,
		"green":     termbox.ColorGreen,
		"yellow":    termbox.ColorYellow,
		"blue":      termbox.ColorBlue,
		"magenta":   termbox.ColorMagenta,
		"cyan":      termbox.ColorCyan,
		"white":     termbox.ColorWhite,
		"bold":      termbox.AttrBold,
		"underline": termbox.AttrUnderline,
		"reverse":   termbox.AttrReverse,
	}

	g := termbox.Attribute(0)
	for _, sv := range s {
		val, ok := values[sv]
		if ok {
			g = g | val
		} else {
			return nil, fmt.Errorf("Unknown value: %v", sv)
		}
	}
	return g, nil
}

type gravityConverter struct{}

func (ac gravityConverter) FromInt(i int) (interface{}, error) {
	return Gravity(i), nil
}
func (ac gravityConverter) FromStrings(s ...string) (interface{}, error) {
	values := map[string]Gravity{
		"left":   Left,
		"right":  Right,
		"bottom": Bottom,
		"top":    Top,
	}

	g := NoGravity
	for _, sv := range s {
		val, ok := values[sv]
		if ok {
			g = g | val
		} else {
			return nil, fmt.Errorf("Unknown value: %v", sv)
		}
	}
	return g, nil
}

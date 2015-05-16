package termui

import (
	"fmt"
	"github.com/boombuler/termui/css"
	"github.com/nsf/termbox-go"
)

var (
	// BackgroundProperty is a css property for the background color.
	BackgroundProperty = NewAttributeProperty("background", css.PropertyBehaviorInherit)
	// ForegroundProperty is a css property for the foreground color
	ForegroundProperty = NewAttributeProperty("color", css.PropertyBehaviorInherit)
	// GravityProperty is a css property for the gravity.
	GravityProperty = NewGravityProperty("gravity", css.PropertyBehaviorDefault)
)

// AttributePropertyDefinition is a typesafe encapsulation of a css property for termbox.Attributes values.
type AttributePropertyDefinition struct {
	*css.Property
}

// NewAttributeProperty creates a new AttributePropertyDefinition
func NewAttributeProperty(name string, b css.PropertyBehavior) *AttributePropertyDefinition {
	att := new(AttributePropertyDefinition)
	att.Property = css.NewProperty(name, b, att)
	return att
}

// FromInt converts an integer to a termbox attribute.
func (a *AttributePropertyDefinition) FromInt(i int) (interface{}, error) {
	return termbox.Attribute(i), nil
}

// FromStrings converts strings to a termbox attribute.
func (a *AttributePropertyDefinition) FromStrings(s ...string) (interface{}, error) {
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

// Get returns the value of the property for the given element.
func (a *AttributePropertyDefinition) Get(e Element) termbox.Attribute {
	style := css.Get(e)
	if val, ok := style[a.Property]; ok {
		if attr, ok := val.(termbox.Attribute); ok {
			return attr
		}
	}
	return termbox.ColorDefault
}

// GravityPropertyDefinition is a typesafe encapsulation of a css property for Gravity values.
type GravityPropertyDefinition struct {
	*css.Property
}

// NewGravityProperty creates a new GravityPropertyDefinition
func NewGravityProperty(name string, b css.PropertyBehavior) *GravityPropertyDefinition {
	gp := new(GravityPropertyDefinition)
	gp.Property = css.NewProperty(name, b, gp)
	return gp
}

// Get returns the value of the property for the given element.
func (gp *GravityPropertyDefinition) Get(e Element) Gravity {
	style := css.Get(e)
	if gravAny, ok := style[gp.Property]; ok {
		if gravity, ok := gravAny.(Gravity); ok {
			return gravity
		}
	}
	return NoGravity
}

// FromInt converts an integer to a gravity value.
func (gp *GravityPropertyDefinition) FromInt(i int) (interface{}, error) {
	return Gravity(i), nil
}

// FromStrings converts strings to a gravity value.
func (gp GravityPropertyDefinition) FromStrings(s ...string) (interface{}, error) {
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

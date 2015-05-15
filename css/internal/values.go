package internal

import (
	"github.com/boombuler/termui/css"
)

type PropertyValue struct {
	Name  string
	Value interface{}
}

type PropertyValues []PropertyValue

type Rule struct {
	Selector css.Selector
	Values   PropertyValues
}

type Rules []Rule

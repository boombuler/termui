package css

import (
	"fmt"
)

type PropertyBehavior int

const (
	PropertyBehaviorDefault PropertyBehavior = 0
	PropertyBehaviorInherit PropertyBehavior = 1 << iota
)

var registeredPropierties map[string]*Property

type Property struct {
	name      string
	behavior  PropertyBehavior
	converter PropertyConverter
}

func (p *Property) Convert(value interface{}) (interface{}, error) {
	if iv, ok := value.(int); ok {
		return p.converter.FromInt(iv)
	} else if sv, ok := value.(string); ok {
		return p.converter.FromStrings(sv)
	} else if svs, ok := value.([]string); ok {
		return p.converter.FromStrings(svs...)
	}
	return nil, fmt.Errorf("Invalid property value: %v", value)
}

type PropertyConverter interface {
	FromInt(i int) (interface{}, error)
	FromStrings(s ...string) (interface{}, error)
}

func NewProperty(name string, behavior PropertyBehavior, conv PropertyConverter) *Property {
	if registeredPropierties == nil {
		registeredPropierties = make(map[string]*Property)
	}
	_, exists := registeredPropierties[name]
	if exists {
		panic("Property " + name + " already exists!")
	}

	prop := &Property{name, behavior, conv}
	registeredPropierties[name] = prop
	return prop
}

func FindProperty(name string) *Property {
	p, ok := registeredPropierties[name]
	if ok {
		return p
	} else {
		return nil
	}
}

func MustFindProperty(name string) *Property {
	p := FindProperty(name)
	if p == nil {
		panic("Property not found: " + name)
	}
	return p
}

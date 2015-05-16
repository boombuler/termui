package css

import (
	"fmt"
)

// PropertyBehavior defines the behavior of a styleable property
type PropertyBehavior int

const (
	// PropertyBehaviorDefault is the default behavior of a property.
	PropertyBehaviorDefault PropertyBehavior = 0
	// PropertyBehaviorInherit can be set to let the style also match all child elements of the element.
	PropertyBehaviorInherit PropertyBehavior = 1 << iota
)

var registeredPropierties map[string]*Property

// Property defines a styleable property.
type Property struct {
	name      string
	behavior  PropertyBehavior
	converter PropertyConverter
}

// Convert calls the converter of the property to cast the value from the style definition to the real type
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

// PropertyConverter is called to convert the value from the css definition to the final type of a property.
type PropertyConverter interface {
	// FromInt casts an integer to the real property value. If the value can not be converted an error should be returned.
	FromInt(i int) (interface{}, error)
	// FromStrings casts one or more strings to the real property value. If the value can not be converted an error should be returned.
	FromStrings(s ...string) (interface{}, error)
}

// NewProperty registers a new property for the styling engine.
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

// FindProperty searches for a registered property and returns the already registered property or nil if no property was found.
func FindProperty(name string) *Property {
	p, ok := registeredPropierties[name]
	if ok {
		return p
	}
	return nil
}

// MustFindProperty searches for a registered property and returns the already registered property. If no property was found the programm panics!
func MustFindProperty(name string) *Property {
	p := FindProperty(name)
	if p == nil {
		panic("Property not found: " + name)
	}
	return p
}

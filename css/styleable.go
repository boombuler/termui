package css

// Styleable defines an interface for styleable elements.
type Styleable interface {
	// ID returns the id of the element.
	ID() string
	// Name returns the name of the element.
	Name() string
	// Classes returns all classes of the element.
	Classes() []string

	// Parent returns the styleable parent of the element or nil if it has no parent.
	Parent() Styleable
}

// ClassMap is a helper to manage classes for elements.
type ClassMap map[string]struct{}

// Add a style to the classmap if it is not already defined.
func (cm ClassMap) Add(s string) {
	cm[s] = struct{}{}
}

// Remove the style from the classmap if it is defined.
func (cm ClassMap) Remove(s string) {
	delete(cm, s)
}

// Classes returns a all currently set classes.
func (cm ClassMap) Classes() []string {
	res := make([]string, 0, len(cm))
	for k := range cm {
		res = append(res, k)
	}
	return res
}

// IDAndClasses is a helper for implementing simple parts of the Styleable interface.
type IDAndClasses struct {
	id      string
	classes ClassMap
}

// AddClass adds a class to the element
func (s *IDAndClasses) AddClass(name string) {
	if s.classes == nil {
		s.classes = make(ClassMap)
	}
	s.classes.Add(name)
}

// RemoveClass removes the class from the element.
func (s *IDAndClasses) RemoveClass(name string) {
	if s.classes != nil {
		s.classes.Remove(name)
	}
}

// Classes returns all classes of the element.
func (s *IDAndClasses) Classes() []string {
	return s.classes.Classes()
}

// SetID sets the hopefully unique id of the element.
func (s *IDAndClasses) SetID(id string) {
	s.id = id
}

// ID returns the currently set id of the element.
func (s *IDAndClasses) ID() string {
	return s.id
}

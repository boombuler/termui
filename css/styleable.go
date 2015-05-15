package css

type Styleable interface {
	Id() string
	Name() string
	Classes() []string

	Parent() Styleable
}

type IdAndClasses struct {
	id      string
	classes map[string]struct{}
}

func (s *IdAndClasses) AddClass(name string) {
	s.classes[name] = struct{}{}
}
func (s *IdAndClasses) RemoveClass(name string) {
	delete(s.classes, name)
}
func (s *IdAndClasses) Classes() []string {
	res := make([]string, 0, len(s.classes))
	for k := range s.classes {
		res = append(res, k)
	}
	return res
}
func (s *IdAndClasses) SetId(id string) {
	s.id = id
}
func (s *IdAndClasses) Id() string {
	return s.id
}

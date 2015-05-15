package css

type Styleable interface {
	Id() string
	Name() string
	Classes() []string

	Parent() Styleable
}

type ClassMap map[string]struct{}

func (cm ClassMap) Add(s string) {
	cm[s] = struct{}{}
}

func (cm ClassMap) Remove(s string) {
	delete(cm, s)
}

func (cm ClassMap) Classes() []string {
	res := make([]string, 0, len(cm))
	for k := range cm {
		res = append(res, k)
	}
	return res
}

type IdAndClasses struct {
	id      string
	classes ClassMap
}

func (s *IdAndClasses) AddClass(name string) {
	if s.classes == nil {
		s.classes = make(ClassMap)
	}
	s.classes.Add(name)
}
func (s *IdAndClasses) RemoveClass(name string) {
	if s.classes != nil {
		s.classes.Remove(name)
	}
}
func (s *IdAndClasses) Classes() []string {
	return s.classes.Classes()
}
func (s *IdAndClasses) SetId(id string) {
	s.id = id
}
func (s *IdAndClasses) Id() string {
	return s.id
}

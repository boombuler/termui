package css

type Styleable interface {
	Id() string
	Name() string
	Classes() []string

	PseudoClasses() []string
	AddPseudoClass(string)
	RemovePseudoClass(string)

	Parent() Styleable
}

type ClassMap map[string]struct{}

func (cm ClassMap) Add(s string) {
	cm[s] = struct{}{}
}

func (cm ClassMap) Remove(s string) {
	delete(cm, s)
}

func (cm ClassMap) List() []string {
	res := make([]string, 0, len(cm))
	for k := range cm {
		res = append(res, k)
	}
	return res
}

type IdAndClasses struct {
	id       string
	pclasses ClassMap
	classes  ClassMap
}

func (s *IdAndClasses) PseudoClasses() []string {
	return s.pclasses.List()
}
func (s *IdAndClasses) AddPseudoClass(name string) {
	if s.pclasses == nil {
		s.pclasses = make(ClassMap)
	}
	s.pclasses.Add(name)
}
func (s *IdAndClasses) RemovePseudoClass(name string) {
	if s.pclasses != nil {
		s.pclasses.Remove(name)
	}
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
	return s.classes.List()
}
func (s *IdAndClasses) SetId(id string) {
	s.id = id
}
func (s *IdAndClasses) Id() string {
	return s.id
}

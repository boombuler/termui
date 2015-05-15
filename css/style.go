package css

type Style map[*Property]interface{}

// Merge combines two styles. the `other` has higher priority.
func (s Style) Merge(other Style) Style {
	res := make(Style)
	if s != nil {
		for k, v := range s {
			res[k] = v
		}
	}
	if other != nil {
		for k, v := range other {
			res[k] = v
		}
	}
	return res
}

func (s Style) Inherit() Style {
	res := make(Style)
	for k, v := range s {
		if k.behavior&PropertyBehaviorInherit != 0 {
			res[k] = v
		}
	}
	return res
}

func (s Style) Value(p *Property) interface{} {
	v, ok := s[p]
	if !ok {
		return struct{}{}
	}
	return v
}

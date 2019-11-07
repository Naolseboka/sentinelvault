package setx

type StringSet map[string]struct{}

func NewString(items ...string) StringSet {
	s := make(StringSet, len(items))
	for _, i := range items {
		s.Add(i)
	}
	return s
}

func (s StringSet) Add(v string) {
	if s == nil {
		return
	}
	s[v] = struct{}{}
}

func (s StringSet) Has(v string) bool {
	_, ok := s[v]
	return ok
}

func (s StringSet) Remove(v string) { delete(s, v) }

func (s StringSet) Len() int { return len(s) }

func (s StringSet) Slice() []string {
	out := make([]string, 0, len(s))
	for k := range s {
		out = append(out, k)
	}
	return out
}

func (s StringSet) Union(o StringSet) StringSet {
	out := NewString()
	for k := range s {
		out.Add(k)
	}
	for k := range o {
		out.Add(k)
	}
	return out
}

func (s StringSet) Intersect(o StringSet) StringSet {
	out := NewString()
	for k := range s {
		if o.Has(k) {
			out.Add(k)
		}
	}
	return out
}

func (s StringSet) Diff(o StringSet) StringSet {
	out := NewString()
	for k := range s {
		if !o.Has(k) {
			out.Add(k)
		}
	}
	return out
}

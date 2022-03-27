package set

type IntSet []int

func (s *IntSet) Add(value int) {
	for _, v := range *s {
		if v == value {
			return
		}
	}

	*s = append(*s, value)
}

func (s *IntSet) ToArray() []int {
	return *s
}

func (s *IntSet) Unite(o IntSet) IntSet {
	var newSet IntSet

	for _, v := range *s {
		newSet.Add(v)
	}
	for _, v := range o.ToArray() {
		newSet.Add(v)
	}

	return newSet
}

func (s *IntSet) Subtract(o IntSet) IntSet {
	var newSet IntSet

	for _, v := range *s {
		if !o.Contains(v) {
			newSet.Add(v)
		}
	}

	return newSet
}

func (s *IntSet) Size() int {
	return len(*s)
}

func (s *IntSet) Equals(o IntSet) bool {
	s1 := s.Unite(o)
	s2 := o.Unite(*s)

	if len(s1.ToArray()) == len(s.ToArray()) && len(o.ToArray()) == len(s2.ToArray()) {
		return true
	}
	return false
}

func (s *IntSet) Contains(v int) bool {
	for _, i := range *s {
		if i == v {
			return true
		}
	}

	return false
}

type ByteSet []byte

func (s *ByteSet) Add(value byte) {
	for _, v := range *s {
		if v == value {
			return
		}
	}

	*s = append(*s, value)
}

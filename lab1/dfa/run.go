package dfa

func (dfa *MinDFA) Run(s string) bool {
	curState := dfa.q0

	for _, c := range s {
		match := false
		for _, t := range dfa.d {
			if t.From == curState && t.Symbol == byte(c) {
				match = true
				curState = t.To
				break
			}
		}
		if !match {
			return false
		}
	}

	matched := false
	for _, f := range dfa.f {
		if curState == f {
			matched = true
			break
		}
	}

	return matched
}

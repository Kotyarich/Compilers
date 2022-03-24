package dfa

import "cc/lab/set"

type DFA struct {
	Q  []*state
	T  set.ByteSet
	D  []tran
	Q0 set.IntSet
	F  []set.IntSet
}

type tran struct {
	state     set.IntSet
	symbol    byte
	destState set.IntSet
}

type state struct {
	value  set.IntSet
	marked bool
}

func (dfa *DFA) hasUnmarked() bool {
	for _, s := range dfa.Q {
		if !s.marked {
			return true
		}
	}

	return false
}

func (dfa *DFA) getUnmarkedPos() int {
	for i, s := range dfa.Q {
		if !s.marked {
			return i
		}
	}

	return -1
}

func isSymbol(s byte) bool {
	if !isUnaryOperation(s) && !isBinaryOperation(s) &&
		s != openBracket && s != closeBracket && s != '#' {
		return true
	}

	return false
}

func (dfa *DFA) getAlphabet(re string) {
	for _, c := range re {
		if isSymbol(byte(c)) {
			dfa.T.Add(byte(c))
		}
	}
}

func (dfa *DFA) hasState(s set.IntSet) bool {
	for i := range dfa.Q {
		if dfa.Q[i].value.Equals(s) {
			return true
		}
	}

	return false
}

func (dfa *DFA) buildF(endPos int) {
	for _, s := range dfa.Q {
		if s.value.Contains(endPos) {
			dfa.F = append(dfa.F, s.value)
		}
	}
}

func BuildDFA(t *tree, re string, followPos map[int]set.IntSet) DFA {
	var dfa DFA
	dfa.getAlphabet(re)

	treeMap := toTreeMap(t)

	dfa.Q0 = t.FirstPos
	dfa.Q = append(dfa.Q, &state{dfa.Q0, false})

	for dfa.hasUnmarked() {
		R := dfa.Q[dfa.getUnmarkedPos()]
		R.marked = true

		for _, symbol := range dfa.T {
			var u set.IntSet
			for _, p := range R.value.ToArray() {
				if treeMap[p].Value[0] == symbol {
					u = u.Unite(followPos[p])
				}
			}

			if !dfa.hasState(u) {
				dfa.Q = append(dfa.Q, &state{u, false})
			}
			dfa.D = append(dfa.D, tran{R.value, symbol, u})
		}
	}

	dfa.buildF(t.Children[1].Pos)

	return dfa
}

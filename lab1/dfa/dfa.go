package dfa

import (
	"cc/lab/set"
)

type tempDFA struct {
	q []*state
	t  set.ByteSet
	d  []tran
	q0 set.IntSet
	f  []set.IntSet
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

func toTreeMap(t *tree) map[int]*tree {
	m := make(map[int]*tree)
	m[t.Pos] = t

	for _, c := range t.Children {
		treeMapRec(c, m)
	}

	return m
}

func treeMapRec(t *tree, m map[int]*tree) {
	m[t.Pos] = t

	for _, c := range t.Children {
		treeMapRec(c, m)
	}
}

func (dfa *tempDFA) hasUnmarked() bool {
	for _, s := range dfa.q {
		if !s.marked {
			return true
		}
	}

	return false
}

func (dfa *tempDFA) getUnmarkedPos() int {
	for i, s := range dfa.q {
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

func (dfa *tempDFA) getAlphabet(re string) {
	for _, c := range re {
		if isSymbol(byte(c)) {
			dfa.t.Add(byte(c))
		}
	}
}

func (dfa *tempDFA) hasState(s set.IntSet) bool {
	for i := range dfa.q {
		if dfa.q[i].value.Equals(s) {
			return true
		}
	}

	return false
}

func (dfa *tempDFA) buildF(endPos int) {
	for _, s := range dfa.q {
		if s.value.Contains(endPos) {
			dfa.f = append(dfa.f, s.value)
		}
	}
}

func BuildDFA(t *tree, re string, followPos map[int]set.IntSet) tempDFA {
	var dfa tempDFA
	dfa.getAlphabet(re)

	treeMap := toTreeMap(t)

	dfa.q0 = t.FirstPos
	dfa.q = append(dfa.q, &state{dfa.q0, false})

	for dfa.hasUnmarked() {
		R := dfa.q[dfa.getUnmarkedPos()]
		R.marked = true

		for _, symbol := range dfa.t {
			var u set.IntSet
			for _, p := range R.value.ToArray() {
				if treeMap[p].Value[0] == symbol {
					u = u.Unite(followPos[p])
				}
			}

			if !dfa.hasState(u) {
				dfa.q = append(dfa.q, &state{u, false})
			}
			dfa.d = append(dfa.d, tran{R.value, symbol, u})
		}
	}

	dfa.buildF(t.Children[1].Pos)

	return dfa
}

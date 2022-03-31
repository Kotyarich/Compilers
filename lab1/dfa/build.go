package dfa

import (
	"cc/lab/set"
	"fmt"
	"strings"
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

func buildDFA(t *tree, re string, followPos map[int]set.IntSet) MinDFA {
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
			for _, p := range R.value {
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

	return toMinDFA(dfa)
}


func preProcessRE(re string) string {
	builder := strings.Builder{}
	builder.WriteString("(")

	for i := 0; i < len(re) - 1; i++ {
		builder.WriteByte(re[i])
		if !isBinaryOperation(re[i]) && re[i] != openBracket &&
			!isBinaryOperation(re[i + 1]) && re[i + 1] != closeBracket && !isUnaryOperation(re[i + 1]) {
			builder.WriteByte('.')
		}
	}

	builder.WriteByte(re[len(re) - 1])
	builder.WriteString(").#")

	return builder.String()
}

func Build(re string, withPrint bool) MinDFA {
	re = preProcessRE(re)

	tree := reToTree(re)
	m := prepareTree(tree)
	DFA := buildDFA(tree, re, m)

	if withPrint {
		fmt.Println("ДКА, полученный из регулярного выражения:")
		DFA.printConfiguration()
	}

	MinimiseOptimal(&DFA)

	if withPrint {
		fmt.Println("Минимизированный ДКА:")
		DFA.printConfiguration()
	}

	return DFA
}

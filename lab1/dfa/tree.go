package dfa

import (
	"cc/lab/set"
	stack2 "cc/lab/stack"
)

const (
	concatOperation      = '.'
	alternativeOperation = '|'
	iterationOperation   = '*'
	openBracket          = '('
	closeBracket         = ')'
)

type tree struct {
	Value    string
	Pos      int
	Nullable bool
	FirstPos set.IntSet
	LastPos  set.IntSet
	Children []*tree
}

func isBinaryOperation(c byte) bool {
	return c == concatOperation || c == alternativeOperation
}

func isUnaryOperation(c byte) bool {
	return c == iterationOperation
}

func oprtPrior(oprt byte) int {
	switch oprt {
	case alternativeOperation:
		return 1
	case concatOperation:
		return 2
	case iterationOperation:
		return 3
	}
	return 0
}

func hangUp(st1 *stack2.Stack, st2 *stack2.Stack, curPos int) {
	value := st1.Pop().(byte)
	p := &tree{Value: string(value), Pos: curPos}

	right := st2.Pop().(*tree)
	if isBinaryOperation(value) {
		left := st2.Pop().(*tree)
		p.Children = append(p.Children, left, right)
	} else {
		p.Children = append(p.Children, right)
	}

	st2.Push(p)
}

func reToTree(expr string) *tree {
	curPos := 0

	st1 := stack2.Stack{}
	st2 := stack2.Stack{}

	l := 0

	p := &tree{}
	l = len(expr)

	for i := 0; i < l; i++ {
		if expr[i] == closeBracket {
			for st1.Top().(byte) != openBracket {
				hangUp(&st1, &st2, curPos)
				curPos++
			}
			st1.Pop()
		} else if expr[i] == openBracket {
			st1.Push(expr[i])
		} else if isBinaryOperation(expr[i]) || isUnaryOperation(expr[i]) {
			for !st1.Empty() && (oprtPrior(st1.Top().(byte)) >= oprtPrior(expr[i])) {
				hangUp(&st1, &st2, curPos)
				curPos++
			}
			st1.Push(expr[i])
		} else {
			p = &tree{Value: string(expr[i]), Pos: curPos}
			curPos++
			st2.Push(p)
		}
	}

	for !st1.Empty() {
		hangUp(&st1, &st2, curPos)
	}

	return st2.Top().(*tree)
}

func prepareTree(t *tree) map[int]set.IntSet {
	m := make(map[int]set.IntSet)
	prepareTreeRecursive(t, m)
	return m
}

func prepareTreeRecursive(t *tree, m map[int]set.IntSet) {
	for _, child := range t.Children {
		prepareTreeRecursive(child, m)
	}

	t.Nullable = nullable(t)
	t.FirstPos = firstPos(t)
	t.LastPos = lastPos(t)
	followPos(t, m)
}

func nullable(t *tree) bool {
	switch t.Value {
	case "e":
		return false
	case string(alternativeOperation):
		for _, child := range t.Children {
			if child.Nullable {
				return true
			}
		}
		return false
	case string(concatOperation):
		for _, child := range t.Children {
			if !child.Nullable {
				return false
			}
		}
		return true
	case string(iterationOperation):
		return true
	default:
		return false
	}
}

func firstPos(t *tree) set.IntSet {
	var s set.IntSet

	switch t.Value {
	case string(alternativeOperation):
		u := t.Children[0]
		v := t.Children[1]
		s = u.FirstPos.Unite(v.FirstPos)
	case string(concatOperation):
		u := t.Children[0]
		v := t.Children[1]
		if u.Nullable {
			s = u.FirstPos.Unite(v.FirstPos)
		} else {
			s = u.FirstPos
		}
	case string(iterationOperation):
		u := t.Children[0]
		s = u.FirstPos
	default:
		s.Add(t.Pos)
	}

	return s
}

func lastPos(t *tree) set.IntSet {
	var s set.IntSet

	switch t.Value {
	case string(alternativeOperation):
		u := t.Children[0]
		v := t.Children[1]
		s = u.LastPos.Unite(v.LastPos)
	case string(concatOperation):
		u := t.Children[0]
		v := t.Children[1]
		if v.Nullable {
			s = u.FirstPos.Unite(v.FirstPos)
		} else {
			s = v.FirstPos
		}
	case string(iterationOperation):
		u := t.Children[0]
		s = u.FirstPos
	default:
		s.Add(t.Pos)
	}

	return s
}

func followPos(t *tree, m map[int]set.IntSet) {
	switch t.Value {
	case string(concatOperation):
		for _, i := range t.Children[0].LastPos.ToArray() {
			curSet := m[i]
			m[i] = curSet.Unite(t.Children[1].FirstPos)
		}
	case string(iterationOperation):
		for _, i := range t.LastPos.ToArray() {
			curSet := m[i]
			m[i] = curSet.Unite(t.FirstPos)
		}
	}
}
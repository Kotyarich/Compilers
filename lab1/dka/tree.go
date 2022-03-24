package dka

import stack2 "cc/lab/stack"

const (
	concatOperation      = '.'
	alternativeOperation = '|'
	iterationOperation   = '*'
	openBracket          = '('
	closeBracket         = ')'
)

type tree struct {
	Value string
	Left  *tree
	Right *tree
}

func isBinaryOperation(c byte) bool {
	return c == concatOperation || c == alternativeOperation
}

//func isUnaryOperation(c rune) bool {
//	return c == iterationOperation
//}
//
//
//func BuildParseTree(re string) *tree {
//	stack := stack2.Stack{}
//	reTree := &tree{}
//	currentTree := reTree
//	stack.Push(currentTree)
//
//	for _, token := range re {
//		if token == openBracket {
//			currentTree.Left = &tree{}
//			stack.Push(currentTree)
//			currentTree = currentTree.Left
//		} else if isBinaryOperation(token) {
//			currentTree.Value = string(token)
//			currentTree.Right = &tree{}
//			stack.Push(currentTree)
//			currentTree = currentTree.Right
//		} else if isUnaryOperation(token) {
//
//		} else if token == closeBracket {
//			currentTree = stack.Pop().(*tree)
//		} else {
//			currentTree.Value = string(token)
//			currentTree = stack.Pop().(*tree)
//		}
//	}
//
//	return reTree
//}

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

func hangUp(st1 *stack2.Stack, st2 *stack2.Stack) {
	p := &tree{Value:string(st1.Pop().(byte))}
	p.Left = st2.Pop().(*tree)
	p.Right = st2.Pop().(*tree)
	st2.Push(p)
}

func ToTree(expr string) *tree {
	st1 := stack2.Stack{}
	st2 := stack2.Stack{}

	l := 0

	p := &tree{}
	l=len(expr)

	for i:=0;i<l;i++ {
		if expr[i]==')' {
			for st1.Top().(byte) !='(' {
				hangUp(&st1, &st2)
			}
			st1.Pop()
		} else if expr[i] == '(' {
			st1.Push(expr[i])
		} else if isBinaryOperation(expr[i]) {
			for !st1.Empty() && (oprtPrior(st1.Top().(byte)) >= oprtPrior(expr[i])) {
				hangUp(&st1, &st2)
			}
			st1.Push(expr[i])
		} else {
			p = &tree{Value: string(expr[i])}
			st2.Push(p)
		}
	}

	for !st1.Empty() {
		hangUp(&st1, &st2)
	}

	return st2.Top().(*tree)
}
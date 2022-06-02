package analyser

import (
	"fmt"
)

const (
	binaryOP int = iota
	unaryOP
	constant
	bracket
)

func opType(op string, expressions []string) int {
	if len(expressions) == 1 && (op == "+" || op == "-") {
		return unaryOP
	}
	if len(expressions) == 2 {
		return binaryOP
	}
	if op == "(" {
		return bracket
	}
	return constant
}

func handleExpression(expressions []string, op string) (string, bool) {
	opType := opType(op, expressions)
	switch opType {
	case binaryOP:
		res := fmt.Sprintf("%s %s %s", expressions[0], expressions[1], op)
		return res, true
	case unaryOP:
		res := fmt.Sprintf("%s %s", expressions[0], op)
		return res, true
	case bracket:
		res := expressions[0]
		return res, true
	}

	return "", false
}

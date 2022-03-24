package main

import (
	"cc/lab/dfa"
	"fmt"
)

func main() {
	re := "(a|b)*.a.b.b"
	re = re + ".#"

	tree := dfa.REToTree(re)
	m := dfa.PrepareTree(tree)
	DFA := dfa.BuildDFA(tree, re, m)
	minDFA := dfa.ToMinDFA(DFA)
	minStates := dfa.Minimise(minDFA)
	
	fmt.Println(minStates)
}

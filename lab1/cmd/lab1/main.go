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
	minStates := dfa.MinimiseOptimal(DFA)
	
	fmt.Println(minStates)
}

package dfa

import (
	"fmt"
)

func (dfa *MinDFA) printConfiguration() {
	dfa.printStates()
	dfa.printTable()
	fmt.Println()
}

func (dfa *MinDFA) printStates() {
	fmt.Println("Q0:", dfa.q0)
	fmt.Print("F: {")
	for i, f := range dfa.f {
		fmt.Print(f)
		if i != len(dfa.f) - 1 {
			fmt.Print(",")
		}
	}
	fmt.Println("}")
}

func (dfa *MinDFA) printTable()  {
	fmt.Println("Таблица переходов:")
	fmt.Print("\t")

	for _, state := range dfa.q {
		fmt.Print(state, "\t")
	}
	fmt.Println()

	for _, state := range dfa.q {
		fmt.Print(state, "\t")
		for _, toState := range dfa.q {
			printed := false
			for _, t := range dfa.d {
				if state == t.From && toState == t.To{
					fmt.Print(string(t.Symbol), "\t")
					printed = true
					break
				}
			}

			if !printed {
				fmt.Print("\t")
			}
		}
		fmt.Println()
	}
}
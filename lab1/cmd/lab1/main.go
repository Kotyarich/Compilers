package main

import (
	"cc/lab/dfa"
	"cc/lab/input"
	"fmt"
	"os"
)

const (
	stdInput  = "-s"
	fileInput = "-f"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("Введите первым аргументом:\n\t-s для ввода из командной строки\n\t-f - из файла 'res.txt'")
		fmt.Println("Введите вторым аргументом регулярное выражение")
	}

	inputType := args[1]
	re := args[2]

	DFA := dfa.Build(re, true)

	var reader input.REReader
	switch inputType {
	case stdInput:
		reader = input.NewStdIREReader()
	case fileInput:
		reader = input.NewFileREReader()
	}

	s, ok := reader.NextRE()
	for ; ok; s, ok = reader.NextRE() {
		matched := DFA.Run(s)
		fmt.Println(s+":", matched)
	}
}

package main

import (
	"fmt"
	"os"

	"lab4/analyser"
	"lab4/reader"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Введите первым аргументом имя файла")
	}

	fileName := args[1]

	r := reader.NewFileReader(fileName)
	a := analyser.NewAnalyser(r)

	expression, ok := a.Analyse()
	if !ok {
		_, _ = fmt.Fprintln(os.Stderr, "error")
	} else {
		fmt.Println(expression)
	}
}

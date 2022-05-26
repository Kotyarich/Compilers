package main

import (
	"fmt"
	"lab3/analyser"
	"lab3/reader"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Введите первым аргументом имя файла")
	}

	fileName := args[1]

	r := reader.NewFileReader(fileName)
	a := analyser.NewAnalyser(r)
	err := a.Parse()
	fmt.Println(err.Error())
}

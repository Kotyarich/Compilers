package main

import (
	"cc/lab/dka"
	"fmt"
)

func main() {
	re := "(a|b).a.#"
	tree := dka.ToTree(re)
	fmt.Println(tree)
}

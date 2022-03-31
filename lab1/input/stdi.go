package input

import (
	"bufio"
	"fmt"
	"os"
)

type StdIREReader struct {
	scanner *bufio.Scanner
}

func NewStdIREReader() *StdIREReader {
	return &StdIREReader{
		scanner: bufio.NewScanner(os.Stdin),
	}
}

func (reader *StdIREReader) NextRE() (string, bool) {
	fmt.Println("Введите строку или \\q для выхода:")

	reader.scanner.Scan()
	re := reader.scanner.Text()

	if re == "\\q" {
		return "", false
	}
	return re, true
}
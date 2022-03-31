package input

import (
	"bufio"
	"log"
	"os"
)

const fileName = "res.txt"

type FileREReader struct {
	file    *os.File
	scanner *bufio.Scanner
}

func NewFileREReader() *FileREReader {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return &FileREReader{
		file:    file,
		scanner: bufio.NewScanner(file),
	}
}

func (reader *FileREReader) NextRE() (string, bool) {
	scanned := reader.scanner.Scan()
	if !scanned {
		_ = reader.file.Close()
		return "", false
	}

	re := reader.scanner.Text()
	return re, true
}

package reader

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type FileReader struct {
	file    *os.File
	scanner *bufio.Reader
}

func NewFileReader(fileName string) *FileReader {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return &FileReader{
		file:    file,
		scanner: bufio.NewReader(file),
	}
}

func (reader *FileReader) NextToken() (string, bool) {
	reader.skipSpaces()

	builder := strings.Builder{}

	for {
		b, err := reader.scanner.ReadByte()
		if err == bufio.ErrFinalToken {
			_ = reader.file.Close()
			return "", false
		}

		builder.WriteByte(b)

		if isToken(builder.String()) {
			break
		}
	}

	return builder.String(), true
}

func (reader *FileReader) UnreadToken(token string) {
	for range token {
		_ = reader.scanner.UnreadByte()
	}
}

func (reader *FileReader) skipSpaces() {
	for {
		b, err := reader.scanner.ReadByte()
		if err == bufio.ErrFinalToken {
			return
		}

		if b == ' ' || b == '\n' {
			continue
		}
		_ = reader.scanner.UnreadByte()
		break
	}
}


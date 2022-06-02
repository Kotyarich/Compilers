package reader

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type FileReader struct {
	file        *os.File
	scanner     *bufio.Reader
	curLine     int64
	curPos      int64
	curTokenLen int64

	closed bool
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

func (reader *FileReader) Next() (string, bool) {
	if reader.closed {
		return "", false
	}
	reader.skipSpaces()

	builder := strings.Builder{}
	buf := make([]byte, 1)
	for {
		reader.curPos++
		reader.curTokenLen++
		_, err := reader.file.Read(buf)
		b := buf[0]
		if err != nil {
			_ = reader.file.Close()
			return "$", true
		}

		builder.WriteByte(b)

		token, ok := reader.isToken(builder.String())
		if ok {
			reader.curTokenLen = 0
			return token, true
		}
	}
}

func (reader *FileReader) UnreadToken(token string) {
	l := int64(len(token))
	reader.curPos -= l
	_, _ = reader.file.Seek(-l, 1)
}

func (reader *FileReader) skipSpaces() {
	buf := make([]byte, 1)
	for {
		_, err := reader.file.Read(buf)
		b := buf[0]
		if err == io.EOF {
			return
		}

		if b == ' ' {
			reader.curPos++
			continue
		}

		if b == 13 {
			reader.curLine++
			reader.curPos = 0
			continue
		}

		if b == 10 {
			continue
		}
		_, _ = reader.file.Seek(-1, 1)

		break
	}
}

func (reader *FileReader) CurPose() (int64, int64) {
	curPos := reader.curPos - reader.curTokenLen
	if curPos < 0 {
		curPos = 0
	}
	return reader.curLine + 1, curPos
}

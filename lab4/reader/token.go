package reader

var tokens map[string]struct{}

func init() {
	tokens = map[string]struct{}{
		"i":  {},
		"n":  {},
		"+":  {},
		"-":  {},
		"*":  {},
		"/":  {},
		"%":  {},
		"<":  {},
		"<=": {},
		"=":  {},
		">=": {},
		">":  {},
		"<>": {},
		"(":  {},
		")":  {},
		";":  {},
		"{":  {},
		"}":  {},
		"^":  {},
	}
}

func (reader *FileReader) isToken(s string) (string, bool) {
	if s == "<" || s == ">" {
		b := make([]byte, 1)
		_, _ = reader.file.Read(b)

		newS := s + string(b)

		_, ok := tokens[newS]
		if ok {
			return newS, ok
		}

		reader.UnreadToken(" ")
	}

	_, ok := tokens[s]
	return s, ok
}

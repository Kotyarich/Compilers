package reader

var tokens map[string]bool

func init() {
	tokens = map[string]bool{
		"<программа>": true,
		"<блок>": true,
		"<список операторов>": true,
		"<хвост>": true,
		"<оператор>": true,
		"id": false,
		"<выражение>": true,
		"<арифметическое выражение>": true,
		"<терм>": true,
		"<знак операции типа сложения>": true,
		"<множитель>": true,
		"<знак операции типа умножения>": true,
		"<первичное выражение>": true,
		"number": false,
		"<знак операции отношения>": true,
		"+": false,
		"-": false,
		"*": false,
		"/": false,
		"%": false,
		"<": false,
		"<=": false,
		"=": false,
		">=": false,
		">": false,
		"<>": false,
		"(": false,
		")": false,
		";": false,
		"{": false,
		"}": false,
		"^": false,
	}
}

func (reader *FileReader) isToken(s string) bool {
	if s == "<" || s == ">" {
		b := make([]byte, 1)
		_, _ = reader.file.Read(b)

		newS := s + string(b)

		_, ok := tokens[newS]
		if ok {
			return ok
		}

		reader.UnreadToken(" ")
	}

	_, ok := tokens[s]
	return ok
}

func tokenType(s string) bool {
	return tokens[s]
}

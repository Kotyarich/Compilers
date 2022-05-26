package reader

var tokens map[string]bool

func init() {
	tokens = map[string]bool{
		"<программа>": true,
		"<блок>": true,
		"<список операторов>": true,
		"<хвост>": true,
		"<оператор>": true,
		"<идентификатор>": false,
		"<выражение>": true,
		"<арифметическое выражение>": true,
		"<терм>": true,
		"<знак операции типа сложения>": true,
		"<множитель>": true,
		"<знак операции типа умножения>": true,
		"<первичное выражение>": true,
		"<число>": false,
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

func isToken(s string) bool {
	_, ok := tokens[s]
	return ok
}

func tokenType(s string) bool {
	return tokens[s]
}

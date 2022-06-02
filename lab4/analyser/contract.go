package analyser

type reader interface {
	Next() (string, bool)
}
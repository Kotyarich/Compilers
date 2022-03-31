package input

type REReader interface {
	NextRE() (string, bool)
}
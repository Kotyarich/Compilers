package analyser

type reader interface {
	NextToken() (string, bool)
	UnreadToken(token string)
	CurPose() (int64, int64)
}
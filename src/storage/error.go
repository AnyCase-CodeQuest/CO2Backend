package storage

type errorStorage struct {
	s string
}

func (e *errorStorage) Error() string {
	return e.s
}

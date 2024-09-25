package app

type Store interface {
	CreateDash(dash Dash) Dash
	ListDashes() []Dash
}

type LocalStore struct {
	dashes []Dash
}

func CreateLocalStore() *LocalStore {
	return &LocalStore{
		dashes: make([]Dash, 0),
	}
}

func (s *LocalStore) CreateDash(dash Dash) Dash {
	s.dashes = append(s.dashes, dash)
	return dash
}

func (s *LocalStore) ListDashes() []Dash {
	return s.dashes
}

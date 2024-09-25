package app

type Store struct {
	dashes []Dash
}

func CreateStore() *Store {
	return &Store{
		dashes: make([]Dash, 0),
	}
}

func (s *Store) CreateDash(dash Dash) Dash {
	s.dashes = append(s.dashes, dash)
	return dash
}

func (s *Store) ListDashes() []Dash {
	return s.dashes
}

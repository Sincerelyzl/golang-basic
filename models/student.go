package models

type Student struct {
	Name  string
	Score int
}

func (s *Student) AddScore(score int) {
	s.Score += score
}

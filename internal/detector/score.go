package detector

type Score struct {
	Total   int
	Reasons []string
}

func (s *Score) add(points int, reason string) {
	s.Total += points
	s.Reasons = append(s.Reasons, reason)
}

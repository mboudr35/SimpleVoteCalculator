package common

// Ballot Maps candidates to a score
type Ballot map[Candidate]int

// GetCandidateScore Get the score of a candidate
func (b Ballot) GetCandidateScore(c Candidate) int {
	return b[c]
}

// SetCandidateScore Set the score of a candidate
func (b Ballot) SetCandidateScore(c Candidate, s int) {
	b[c] = s
}

// NewBallot Create a new ballot with a given number of candidates
func NewBallot(candidateCount int) Ballot {
	return make(Ballot, candidateCount)
}

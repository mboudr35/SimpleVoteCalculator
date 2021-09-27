package common

// Candidate Represents a candidate with an ID number and name
type Candidate struct {
	id   int
	name string
}

// GetId Returns the candidate's ID number
func (c Candidate) GetId() int {
	return c.id
}

// GetName Returns the candidate's name
func (c Candidate) GetName() string {
	return c.name
}

func (c Candidate) String() string {
	return c.name
}

// NewCandidate Creates a new candidate with a given ID number and name
func NewCandidate(id int, name string) Candidate {
	return Candidate{
		id:   id,
		name: name,
	}
}

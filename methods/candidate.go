package methods

// Candidate Represents a candidate with an ID number and name
type Candidate struct {
	id   int
	name string
}

func (c Candidate) GetId() int {
	return c.id
}

func (c Candidate) GetName() string {
	return c.name
}

func (c Candidate) String() string {
	return c.name
}

func NewCandidate(id int, name string) Candidate {
	return Candidate{
		id:   id,
		name: name,
	}
}

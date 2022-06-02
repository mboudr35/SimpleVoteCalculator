package common

// Candidate Represents a candidate with an ID number and name
type Candidate struct {
	id   int
	name string
	rank int
}

// GetId Returns the candidate's ID number
func (c *Candidate) GetId() int {
	return c.id
}

// GetName Returns the candidate's name
func (c *Candidate) GetName() string {
	return c.name
}

func (c *Candidate) GetRank() int {
	return c.rank
}

func (c *Candidate) SetRank(rank int) {
	c.rank = rank
}

func (c *Candidate) IncrRank(incr int) {
	c.rank += incr
}

func (c *Candidate) String() string {
	return c.name
}

// NewCandidate Creates a new candidate with a given ID number and name
func NewCandidate(id int, name string) Candidate {
	return Candidate{
		id:   id,
		name: name,
		rank: 0,
	}
}

type Candidates []Candidate

func (a Candidates) Len() int {
	return len(a)
}

func (a Candidates) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Candidates) Less(i, j int) bool {
	return a[i].rank > a[j].rank
}

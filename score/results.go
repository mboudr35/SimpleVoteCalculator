package score

import (
	"simplevotecalculator/common"
	"sort"
)

// Compute Returns a list of candidates ordered by score
func Compute(candidates common.Candidates, ballots []common.Ballot) {
	for i, c := range candidates {
		for _, b := range ballots {
			candidates[i].IncrRank(b[c])
		}
	}
	sort.Sort(candidates)
}

package schulze

import (
	"simplevotecalculator/common"
	"sort"
)

func Compute(candidates common.Candidates, ballots []common.Ballot) {
	ps := GetPathStrengthMatrix(candidates, GetPreferenceMatrix(candidates, ballots))
	for i, c1 := range candidates {
		for _, c2 := range candidates {
			if c1 != c2 && ps[c1][c2] > ps[c2][c1] {
				candidates[i].IncrRank(1)
			}
		}
	}
	sort.Sort(candidates)
}

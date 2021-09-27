package schulze

import "ctf.mcgill.ca/internal/election/common"

func GetPreferenceMatrix(candidates []common.Candidate, ballots []common.Ballot) CandidateMatrix {
	result := NewCandidateMatrix(candidates)
	for _, c1 := range candidates {
		// Populate the preference matrix
		for _, c2 := range candidates {
			if c1 != c2 {
				for _, b := range ballots {
					// The smaller the score the higher the rank (ie: rank of 1 is best, rank of C is worst)
					if b.GetCandidateScore(c1) < b.GetCandidateScore(c2) {
						result[c1][c2]++
					}
				}
			}
		}
	}
	return result
}

package schulze

import "ctf.mcgill.ca/internal/election/common"

type CandidateMatrix map[common.Candidate]map[common.Candidate]int

func NewCandidateMatrix(candidates []common.Candidate) CandidateMatrix {
	C := len(candidates)
	result := make(CandidateMatrix, C)
	for _, c := range candidates {
		result[c] = make(map[common.Candidate]int, C)
	}
	return result
}

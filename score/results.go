package score

import (
	"ctf.mcgill.ca/internal/election/common"
	"sort"
)

// GetResult Returns a list of candidates ordered by score
func GetResult(candidates []common.Candidate, ballots []common.Ballot) List {
	// Can be adapted to approval by setting approve = 1, neutral = 0, reject = -1.  Winners will have positive score.
	candidateScores := make(List, len(candidates))
	for cid, cval := range candidates {
		candidateScores[cid] = ListElem{
			value: cval,
			score: 0,
			index: cid,
		}
		// Candidate's final score is sum of ballot scores
		for _, bv := range ballots {
			candidateScores[cid].score += bv.GetCandidateScore(cval)
		}
	}
	sort.Sort(candidateScores)
	return candidateScores
}

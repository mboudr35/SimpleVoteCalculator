package methods

import (
	"fmt"
	"sort"
)

type ScoreListElem struct {
	value Candidate
	score int
	index int
}

func (sle ScoreListElem) GetValue() Candidate {
	return sle.value
}

func (sle ScoreListElem) GetScore() int {
	return sle.score
}

func (sle ScoreListElem) GetIndex() int {
	return sle.index
}

func (sle ScoreListElem) String() string {
	return fmt.Sprintf("%s:%d", sle.value, sle.score)
}

type ScoreList []ScoreListElem

func (sl ScoreList) Len() int {
	return len(sl)
}

func (sl ScoreList) Less(i, j int) bool {
	// Highest score first
	return sl[i].score > sl[j].score
}

func (sl ScoreList) Swap(i, j int) {
	sl[i], sl[j] = sl[j], sl[i]
	sl[i].index = i
	sl[j].index = j
}

// GetScoreResult Returns a list of candidates ordered by score
func GetScoreResult(candidates []Candidate, ballots []Ballot) ScoreList {
	candidateScores := make(ScoreList, len(candidates))
	for cid, cval := range candidates {
		candidateScores[cid] = ScoreListElem{
			value: cval,
			score: 0,
			index: cid,
		}
		for _, bv := range ballots {
			candidateScores[cid].score += bv.GetCandidateScore(cval)
		}
	}
	sort.Sort(candidateScores)
	return candidateScores
}

package score

import (
	"ctf.mcgill.ca/internal/election/common"
	"fmt"
)

type ListElem struct {
	value common.Candidate
	score int
	index int
}

func (sle ListElem) GetValue() common.Candidate {
	return sle.value
}

func (sle ListElem) GetScore() int {
	return sle.score
}

func (sle ListElem) GetIndex() int {
	return sle.index
}

func (sle ListElem) String() string {
	return sle.value.String()
}

type List []ListElem

func (sl List) Len() int {
	return len(sl)
}

func (sl List) Less(i, j int) bool {
	// Highest score first
	return sl[i].score > sl[j].score
}

func (sl List) Swap(i, j int) {
	sl[i], sl[j] = sl[j], sl[i]
	sl[i].index = i
	sl[j].index = j
}

func (sl List) String() string {
	var outer [][]common.Candidate
	var inner []common.Candidate
	visited := false
	lastScore := 0
	for _, sle := range sl {
		if visited {
			if lastScore == sle.score {
				inner = append(inner, sle.value)
			} else {
				outer = append(outer, inner)
				inner = []common.Candidate{sle.value}
				lastScore = sle.score
			}
		} else {
			lastScore = sle.score
			inner = append(inner, sle.value)
			visited = true
		}
	}
	outer = append(outer, inner)
	return fmt.Sprint(outer)
}

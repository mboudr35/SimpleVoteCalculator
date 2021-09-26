package methods

import (
	"math"
)

// WinMap A mapping of candidates to their win status
type WinMap map[Candidate]bool

// GetWinners Returns all candidates that have a true win status
func (wm WinMap) GetWinners() []Candidate {
	var resultv []Candidate
	for c, b := range wm {
		if b {
			resultv = append(resultv, c)
		}
	}
	return resultv
}

// GetSchulzePathStrength Calculates the strength of the strongest paths for each pair of candidates
func GetSchulzePathStrength(candidates []Candidate, ballots []Ballot) map[Candidate]map[Candidate]int {
	C := len(candidates)
	// C * C matrix where prefs[a][b] = number of ballots strictly preferring a to b
	prefs := make(map[Candidate]map[Candidate]int, C)
	// C * C matrix where pathStrength[a][b] = strength of the strongest path from a to b
	pathStrength := make(map[Candidate]map[Candidate]int, C)

	// Initialize the structures
	for _, c1 := range candidates {
		// Nested maps need initialization
		prefs[c1] = make(map[Candidate]int, C)
		pathStrength[c1] = make(map[Candidate]int, C)
		// Populate the preference matrix
		for _, c2 := range candidates {
			if c1 != c2 {
				for _, b := range ballots {
					// The smaller the score the higher the rank (ie: rank of 1 is best, rank of C is worst)
					if b.GetCandidateScore(c1) < b.GetCandidateScore(c2) {
						prefs[c1][c2] += 1
					}
				}
			}
		}
	}
	// Path strength
	// Initialize to preference score
	for _, c1 := range candidates {
		for _, c2 := range candidates {
			if c1 != c2 && prefs[c1][c2] > prefs[c2][c1] {
				pathStrength[c1][c2] = prefs[c1][c2]
			}
		}
	}
	// Find the best strength
	for _, c1 := range candidates {
		for _, c2 := range candidates {
			if c1 != c2 {
				for _, c3 := range candidates {
					if c1 != c3 && c2 != c3 {
						pathStrength[c2][c3] = int(math.Max(float64(pathStrength[c2][c3]), math.Min(float64(pathStrength[c2][c1]), float64(pathStrength[c1][c3]))))
					}
				}
			}
		}
	}
	return pathStrength
}

// GetSchulzeWinners Determines the winners of a Schulze iteration
func GetSchulzeWinners(candidates []Candidate, pathStrength map[Candidate]map[Candidate]int) []Candidate {
	winMap := make(WinMap, len(candidates))
	// Candidates to wins map & relation
	for _, c1 := range candidates {
		winMap[c1] = true
		for _, c2 := range candidates {
			if c1 != c2 && pathStrength[c1][c2] < pathStrength[c2][c1] {
				winMap[c1] = false
			}
		}
	}
	return winMap.GetWinners()
}

// GetSchulzeResult Calculates the sorted ranking of candidates in a Schulze methods
func GetSchulzeResult(candidates []Candidate, pathStrength map[Candidate]map[Candidate]int) [][]Candidate {
	var result [][]Candidate
	for len(candidates) > 0 {
		var step []Candidate
		for _, winner := range GetSchulzeWinners(candidates, pathStrength) {
			step = append(step, winner)
			/*delete(pathStrength, winner)
			for _, m := range pathStrength {
				delete(m, winner)
			}*/
			index := -1
			for i, c := range candidates {
				if winner == c {
					index = i
					break
				}
			}
			copy(candidates[index:], candidates[index+1:])
			candidates = candidates[:len(candidates)-1]
		}
		result = append(result, step)
	}
	return result
}

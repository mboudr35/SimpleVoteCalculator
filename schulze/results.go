package schulze

import (
	"ctf.mcgill.ca/internal/election/common"
)

// GetWinners Determines the winners of a Schulze iteration
func GetWinners(candidates []common.Candidate, pathStrength CandidateMatrix) []common.Candidate {
	winMap := make(WinMap, len(candidates))
	// Candidates to wins map
	for _, c1 := range candidates {
		// Assume you're a winner until proven otherwise
		winMap[c1] = true
		for _, c2 := range candidates {
			if c1 != c2 && pathStrength[c1][c2] < pathStrength[c2][c1] {
				winMap[c1] = false
			}
		}
	}
	return winMap.GetWinners()
}

// GetResult Calculates the sorted ranking of candidates in a Schulze election
func GetResult(candidates []common.Candidate, ballots []common.Ballot) [][]common.Candidate {
	prefs := GetPreferenceMatrix(candidates, ballots)
	pathStrength := GetPathStrengthMatrix(candidates, prefs)
	var result [][]common.Candidate
	for len(candidates) > 0 {
		var step []common.Candidate
		// May have ties, put them together in subarray
		for _, winner := range GetWinners(candidates, pathStrength) {
			step = append(step, winner)
			index := -1
			// Find the winner's index
			for i, c := range candidates {
				if winner == c {
					index = i
					break
				}
			}
			// Remove the current winner, will reiterate on remaining candidates to get next best
			copy(candidates[index:], candidates[index+1:])
			candidates = candidates[:len(candidates)-1]
		}
		result = append(result, step)
	}
	return result
}

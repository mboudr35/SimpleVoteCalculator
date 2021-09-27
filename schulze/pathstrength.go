package schulze

import (
	"ctf.mcgill.ca/internal/election/common"
	"math"
)

// GetPathStrengthMatrix Calculates the strength of the strongest paths for each pair of candidates
func GetPathStrengthMatrix(candidates []common.Candidate, prefs CandidateMatrix) CandidateMatrix {
	// C * C matrix where result[a][b] = strength of the strongest path from a to b
	result := NewCandidateMatrix(candidates)

	// Path strength
	// Initialize to preference score
	for _, c1 := range candidates {
		for _, c2 := range candidates {
			if c1 != c2 && prefs[c1][c2] > prefs[c2][c1] {
				result[c1][c2] = prefs[c1][c2]
			}
		}
	}
	// Find the best strength (modified Floyd-Warshall)
	for _, c1 := range candidates {
		for _, c2 := range candidates {
			if c1 != c2 {
				for _, c3 := range candidates {
					if c1 != c3 && c2 != c3 {
						result[c2][c3] = int(math.Max(float64(result[c2][c3]), math.Min(float64(result[c2][c1]), float64(result[c1][c3]))))
					}
				}
			}
		}
	}
	return result
}

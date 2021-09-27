package schulze

import "ctf.mcgill.ca/internal/election/common"

// WinMap A mapping of candidates to their win status
type WinMap map[common.Candidate]bool

// GetWinners Returns all candidates that have a true win status
func (wm WinMap) GetWinners() []common.Candidate {
	var resultv []common.Candidate
	for c, b := range wm {
		if b {
			resultv = append(resultv, c)
		}
	}
	return resultv
}

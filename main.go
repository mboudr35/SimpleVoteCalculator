package main

import (
	"bufio"
	"ctf.mcgill.ca/internal/voting/methods"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var schulze bool
var score bool

func init() {
	flag.BoolVar(&schulze, "schulze", false, "Use the Schulze method")
	flag.BoolVar(&score, "score", false, "Use the score method")
}

func main() {
	flag.Parse()
	csvRead := csv.NewReader(bufio.NewReader(os.Stdin))
	header, err := csvRead.Read()
	if err != nil {
		log.Fatalf("Failed to read header: %v\n", err)
	}
	candidates := make([]methods.Candidate, len(header))
	// CSV header is candidate list
	for column, name := range header {
		candidates[column] = methods.NewCandidate(column, name)
	}
	ballotsIn, err := csvRead.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read ballot table: %v\n", err)
	}
	// Remaining CSV lines are ballots, where each column corresponds to that candidate's score
	ballots := make([]methods.Ballot, len(ballotsIn))
	for ballotIndex, ballotValues := range ballotsIn {
		ballots[ballotIndex] = methods.NewBallot(len(candidates))
		for column, scoreString := range ballotValues {
			score, err := strconv.Atoi(scoreString)
			if err != nil {
				log.Fatalf("Expected integer, got '%s'\n", scoreString)
			}
			ballots[ballotIndex].SetCandidateScore(candidates[column], score)
		}
	}
	// Calculate the results!
	if schulze {
		fmt.Println(methods.GetSchulzeResult(candidates, methods.GetSchulzePathStrength(candidates, ballots)))
	}
	if score {
		fmt.Println(methods.GetScoreResult(candidates, ballots))
	}
}

package main

import (
	"bufio"
	"ctf.mcgill.ca/internal/election/common"
	"ctf.mcgill.ca/internal/election/schulze"
	"ctf.mcgill.ca/internal/election/score"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var schulzeFlag bool
var scoreFlag bool

func usage() {
	fmt.Fprintln(flag.CommandLine.Output(), "Takes a list of comma-separated values (CSV) from standard in.")
	fmt.Fprintln(flag.CommandLine.Output(), "The first line (header) is the list of candidates.")
	fmt.Fprintln(flag.CommandLine.Output(), "The following lines are ballots, where column entries correspond to the candidate's scoreFlag from the ballot.")
	fmt.Fprintln(flag.CommandLine.Output(), "Flags:")
	flag.PrintDefaults()
}

func init() {
	flag.BoolVar(&schulzeFlag, "schulze", false, "Use the Schulze method.  Lower ballot column entries yield a better rank for the corresponding candidate.")
	flag.BoolVar(&scoreFlag, "score", false, "Use the score method.  Higher ballot column entries yield a better scoreFlag for the corresponding candidate.")
	flag.Usage = usage
}

func main() {
	flag.Parse()
	csvRead := csv.NewReader(bufio.NewReader(os.Stdin))
	header, err := csvRead.Read()
	if err != nil {
		log.Fatalf("Failed to read header: %v\n", err)
	}
	candidates := make([]common.Candidate, len(header))
	// CSV header is candidate list
	for column, name := range header {
		candidates[column] = common.NewCandidate(column, name)
	}
	ballotsIn, err := csvRead.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read ballot table: %v\n", err)
	}
	// Remaining CSV lines are ballots, where each column corresponds to that candidate's scoreFlag
	ballots := make([]common.Ballot, len(ballotsIn))
	for ballotIndex, ballotValues := range ballotsIn {
		ballots[ballotIndex] = common.NewBallot(len(candidates))
		for column, scoreString := range ballotValues {
			score, err := strconv.Atoi(scoreString)
			if err != nil {
				log.Fatalf("Expected integer, got '%s'\n", scoreString)
			}
			ballots[ballotIndex].SetCandidateScore(candidates[column], score)
		}
	}
	// Calculate the results!
	if schulzeFlag {
		fmt.Println(schulze.GetResult(candidates, ballots))
	}
	if scoreFlag {
		fmt.Println(score.GetResult(candidates, ballots))
	}
}

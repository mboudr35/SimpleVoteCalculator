package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"simplevotecalculator/common"
	"simplevotecalculator/schulze"
	"simplevotecalculator/score"
	"strconv"
)

var schulzeFlag bool
var scoreFlag bool

func usage() {
	fmt.Fprintln(flag.CommandLine.Output(), "Takes a list of comma-separated values (CSV) from standard in.")
	fmt.Fprintln(flag.CommandLine.Output(), "The first line (header) is the list of candidates.")
	fmt.Fprintln(flag.CommandLine.Output(), "The following lines are ballots, where column entries correspond to the candidate's score from the ballot.")
	fmt.Fprintln(flag.CommandLine.Output(), "Flags:")
	flag.PrintDefaults()
}

func init() {
	flag.BoolVar(&schulzeFlag, "schulze", false, "Use the Schulze method.  Lower ballot column entries yield a better rank for the corresponding candidate.")
	flag.BoolVar(&scoreFlag, "score", false, "Use the score method.  Higher ballot column entries yield a better score for the corresponding candidate.")
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
	// Remaining CSV lines are ballots, where each column corresponds to that candidate's score
	ballots := make([]common.Ballot, len(ballotsIn))
	for ballotIndex, ballotValues := range ballotsIn {
		ballots[ballotIndex] = common.NewBallot(len(candidates))
		for column, scoreString := range ballotValues {
			points, err := strconv.Atoi(scoreString)
			if err != nil {
				log.Fatalf("Expected integer, got '%s'\n", scoreString)
			}
			ballots[ballotIndex].SetCandidateScore(candidates[column], points)
		}
	}
	var computer func(cs common.Candidates, bs []common.Ballot)
	// Calculate the results!
	if schulzeFlag {
		computer = schulze.Compute
	} else if scoreFlag {
		computer = score.Compute
	} else {
		log.Println("No method specified")
		flag.Usage()
		os.Exit(1)
	}
	computer(candidates, ballots)
	csvWrite := csv.NewWriter(os.Stdout)
	outHeader := make([]string, len(candidates))
	outBody := make([]string, len(candidates))
	for i, c := range candidates {
		outHeader[i] = c.GetName()
		outBody[i] = strconv.Itoa(c.GetRank())
	}
	if err := csvWrite.Write(outHeader); err != nil {
		log.Fatalf("Failed to write output CSV header: %v\n", err)
	}
	if err := csvWrite.Write(outBody); err != nil {
		log.Fatalf("Failed to write output CSV body: %v\n", err)
	}
	csvWrite.Flush()
}

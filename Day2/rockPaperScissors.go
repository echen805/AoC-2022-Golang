package main

import (
	"os"
	"strings"
	"log"
	"io"
	"fmt"
)

// score returned for choosing a symbol
var symbolScore = map[string]int {
	"rock": 1,
	"paper": 2,
	"scissors": 3,
}

// score returned given an outcome
var outcomeScore = map[string]int{
	"lost": 0,
	"draw": 3,
	"won": 6,
}

// decode X, Y, Z to lost, draw, won
var decodeOutcome = map[string]string {
	"X": "lost",
	"Y": "draw",
	"Z": "won",
}

// Return the index to search two dim array given "rock, paper, scissors"
var symbolIndex = map[string]int {
	"rock": 0,
	"paper": 1,
	"scissors": 2,
}

// Returns score given t, o, s
var resultsScoreMap = map[string]int{
	"t": 3,
	"o": 0,
	"s": 6,
}

// Lookup symbol, returns rock, paper, scissors
var symbolDict = map[string]string {
	"A": "rock",
	"B": "paper",
	"C": "scissors",
	"X": "rock",
	"Y": "paper",
	"Z": "scissors",
}

// Correct response symbol given XYZ and opp symbol
var responseIndex = map[string]int {
	"X": 0,
	"Y": 1,
	"Z": 2,
}

// Returns the symbol: A, B, C
var myOption = [3][3]string {
	{"C", "A", "B"},
	{"A", "B", "C"},
	{"B", "C", "A"},
}

/*
Results Matrix
t == tie
o == opponent wins
s == strategy wins

    X  Y  Z
A { t  s  o }
B { o  t  s }
C { s  o  t }
*/

// two dimensional array of results 
var results = [3][3] string {
	{"t", "s", "o"},
	{"o", "t", "s"},
	{"s", "o", "t"},
}

func getMatchResults(opponent string, strat string) int {
	strResult := results[symbolIndex[symbolDict[opponent]]][symbolIndex[symbolDict[strat]]]
	return resultsScoreMap[strResult]
}

func calculateRoundScore(opponent string, mine string) int {
	matchResult := getMatchResults(opponent, mine)

	choiceScore := symbolScore[symbolDict[mine]]

	return matchResult + choiceScore
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func calculateCorrectScore(opponent string, strat string) int {
	
	symbolIndexOpp := symbolIndex[symbolDict[opponent]]
	symbolIndexStrat := responseIndex[strat]
	myChoice := symbolDict[myOption[symbolIndexOpp][symbolIndexStrat]]
	choiceScore := symbolScore[myChoice]
	outcome := decodeOutcome[strat]
	matchResult := outcomeScore[outcome]
	
	// fmt.Printf("\nopp choice: %s, strat decision: %s, my choice: %s\n", symbolDict[opponent], outcome, myChoice)

	// fmt.Printf("choice score: %v, matchResult: %v\n", choiceScore, matchResult)
	return choiceScore + matchResult
}

func main() {
	f, err := os.Open("./day2_input.txt")
	check(err)
	// Run after main is done running due to defer
	defer f.Close()
	rawBytes, err := io.ReadAll(f)
	check(err)

	lines := strings.Split(string(rawBytes), "\n")
	optimumScore := 0
	correctScore := 0
	for _, line:= range lines {
		splitRound := strings.Split(line, " ")
		opponentOption := splitRound[0]
		strategyOption := splitRound[1]
		
		optimumScore += calculateRoundScore(opponentOption, strategyOption)

		correctScore += calculateCorrectScore(opponentOption, strategyOption)
	}
	fmt.Println("\nWinning every round score: ", optimumScore)
	fmt.Println("Correct strategy score: ", correctScore)
}
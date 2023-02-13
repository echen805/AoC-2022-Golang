package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// If there's an error, abort, print error and goroutine traces
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func getMaxCalories(file string) {
	// Open the file as a variable
	f, err := os.Open(file)
	check(err)
	// Run after main is done running due to defer
	defer f.Close()
	rawBytes, err := io.ReadAll(f)
	check(err)

	currCalories := 0
	lines := strings.Split(string(rawBytes), "\n")
	elvesCal := make([]int, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			elvesCal = append(elvesCal, currCalories)
			currCalories = 0
		} else {
			intLine, err := strconv.Atoi(line)
			check(err)
			currCalories += intLine
		}
	}
	// Sort in descending order
	sort.Slice(elvesCal, func(i, j int) bool {
		return elvesCal[i] > elvesCal[j]
	})

	fmt.Println("Max calories is", elvesCal[0])
	fmt.Println("Top 3 calories", elvesCal[0]+elvesCal[1]+elvesCal[2])
}

func main() {
	getMaxCalories("./input.txt")
}

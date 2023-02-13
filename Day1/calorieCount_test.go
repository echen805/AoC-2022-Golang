package main

import (
	"log"
	"os"
	"testing"
	// "github.com/echen805/aoc-2022-golang/calorieCount"
)

func TestGetMaxCalories(t *testing.T) {
	f, err := os.Open("./input.txt")
	want := 600
	if got := calorieCount(); got != want {
		t.Errorf("calorieCount() = %q, want %q", got, want)
	}

	// want := "Hello, world."
	// if got := Hello(); got != want {
	// 	t.Errorf("Hello() = %q, want %q", got, want)
	// }
}

package main

import (
	"testing"

	"aoc2025/internal/aoc"
)

func TestPart1(t *testing.T) {
	moves := parse(aoc.ReadInput("sample.txt"))
	got := part1(moves)
	want := 3 // dial lands on 0 three times in the example
	if got != want {
		t.Errorf("part1 = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	t.Skip("part 2 not unlocked yet")
}

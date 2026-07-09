package main

import (
	"testing"

	"aoc2025/internal/aoc"
)

func TestPart1(t *testing.T) {
	lines := aoc.ReadLines("sample.txt")
	got := part1(lines)
	want := 0 // TODO: expected answer for the sample
	if got != want {
		t.Errorf("part1 = %d, want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	lines := aoc.ReadLines("sample.txt")
	got := part2(lines)
	want := 0 // TODO: expected answer for the sample
	if got != want {
		t.Errorf("part2 = %d, want %d", got, want)
	}
}

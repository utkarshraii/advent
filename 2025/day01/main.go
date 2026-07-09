package main

import (
	"flag"
	"fmt"
	"regexp"

	"aoc2025/internal/aoc"
)

// The dial shows 0..99 in a circle and starts at 50.
// L rotates toward lower numbers (0 wraps to 99), R toward higher (99 wraps to 0).
const dialSize = 100
const start = 50

type move struct {
	dir  byte // 'L' or 'R'
	dist int
}

// moveRe pulls out every "L68" / "R14" token, so the input can be separated by
// commas, spaces, or newlines — we don't care.
var moveRe = regexp.MustCompile(`([LR])(\d+)`)

func parse(s string) []move {
	matches := moveRe.FindAllStringSubmatch(s, -1)
	moves := make([]move, len(matches))
	for i, m := range matches {
		moves[i] = move{dir: m[1][0], dist: aoc.MustAtoi(m[2])}
	}
	return moves
}

// part1: count how many times the dial points at 0 after a rotation.
func part1(moves []move) int {
	pos := start
	count := 0
	for _, mv := range moves {
		if mv.dir == 'L' {
			pos -= mv.dist
		} else {
			pos += mv.dist
		}
		pos = ((pos % dialSize) + dialSize) % dialSize // wrap into 0..99
		if pos == 0 {
			count++
		}
	}
	return count
}

func part2(moves []move) int {
	// TODO: unlocks after part 1 is submitted.
	return 0
}

func main() {
	sample := flag.Bool("sample", false, "use sample.txt instead of input.txt")
	flag.Parse()

	file := "input.txt"
	if *sample {
		file = "sample.txt"
	}
	moves := parse(aoc.ReadInput(file))

	fmt.Println("Part 1:", part1(moves))
	fmt.Println("Part 2:", part2(moves))
}

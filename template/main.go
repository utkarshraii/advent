package main

import (
	"flag"
	"fmt"

	"aoc2025/internal/aoc"
)

func main() {
	sample := flag.Bool("sample", false, "use sample.txt instead of input.txt")
	flag.Parse()

	file := "input.txt"
	if *sample {
		file = "sample.txt"
	}
	lines := aoc.ReadLines(file)

	fmt.Println("Part 1:", part1(lines))
	fmt.Println("Part 2:", part2(lines))
}

func part1(lines []string) int {
	return 0
}

func part2(lines []string) int {
	return 0
}

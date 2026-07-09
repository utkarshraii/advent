// Package aoc holds small helpers shared across every Advent of Code day.
package aoc

import (
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

// ReadInput reads a file located next to the source file that calls it, so a
// day can do aoc.ReadInput("input.txt") and it works regardless of the current
// working directory (e.g. `go run ./2025/day01` from the repo root). It returns
// the file contents with trailing newlines trimmed. It panics on error, which
// is fine for throwaway puzzle solutions.
func ReadInput(name string) string {
	_, caller, _, ok := runtime.Caller(1)
	if !ok {
		panic("aoc: could not determine caller for " + name)
	}
	path := filepath.Join(filepath.Dir(caller), name)
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.TrimRight(string(b), "\r\n")
}

// ReadLines is ReadInput split into lines.
func ReadLines(name string) []string {
	_, caller, _, ok := runtime.Caller(1)
	if !ok {
		panic("aoc: could not determine caller for " + name)
	}
	path := filepath.Join(filepath.Dir(caller), name)
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimRight(string(b), "\r\n"), "\n")
}

var intRe = regexp.MustCompile(`-?\d+`)

// Ints extracts every (optionally negative) integer from s, ignoring any other
// characters. Handy for parsing lines like "x=12, y=-3".
func Ints(s string) []int {
	matches := intRe.FindAllString(s, -1)
	out := make([]int, len(matches))
	for i, m := range matches {
		out[i] = MustAtoi(m)
	}
	return out
}

// MustAtoi parses an int or panics.
func MustAtoi(s string) int {
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return n
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Sum returns the sum of xs.
func Sum(xs []int) int {
	total := 0
	for _, x := range xs {
		total += x
	}
	return total
}

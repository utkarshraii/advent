#!/usr/bin/env bash
# Scaffold a new Advent of Code day from the template.
# Usage: scripts/new.sh <day> [year]
#   scripts/new.sh 1        -> 2025/day01
#   scripts/new.sh 7 2024   -> 2024/day07
set -euo pipefail

if [[ $# -lt 1 ]]; then
	echo "usage: $0 <day> [year]" >&2
	exit 1
fi

day=$(printf "%02d" "$1")
year="${2:-2025}"
root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
dir="$root/$year/day$day"

if [[ -d "$dir" ]]; then
	echo "already exists: $dir" >&2
	exit 1
fi

mkdir -p "$dir"
cp "$root/template/main.go"      "$dir/main.go"
cp "$root/template/main_test.go" "$dir/main_test.go"
cp "$root/template/sample.txt"   "$dir/sample.txt"
touch "$dir/input.txt"   # gitignored; paste your puzzle input here

echo "created $dir"
echo "  1. paste your puzzle input into $year/day$day/input.txt"
echo "  2. paste the example into      $year/day$day/sample.txt"
echo "  3. run:  go run ./$year/day$day"
echo "     test: go test ./$year/day$day"

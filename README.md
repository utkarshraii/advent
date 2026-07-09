# Advent of Code 🎄

My [Advent of Code](https://adventofcode.com/) solutions, in Go.

## Layout

```
2025/day01/        one self-contained program per day
  main.go            part1() and part2()
  main_test.go       runs against sample.txt
  sample.txt         the example from the puzzle (committed)
  input.txt          your personal puzzle input (gitignored)
internal/aoc/      shared helpers (ReadLines, Ints, Abs, ...)
template/          source the scaffolder copies for a new day
scripts/new.sh     scaffolder
```

## Workflow

```bash
make new DAY=1          # scaffold 2025/day01
# paste your puzzle input into 2025/day01/input.txt
# paste the example into      2025/day01/sample.txt
make sample DAY=1       # run against the example while developing
make test DAY=1         # check part1/part2 against the sample
make run DAY=1          # run against your real input
```

For a previous year: `make new DAY=1 YEAR=2024`.

Puzzle inputs are gitignored on purpose — AoC asks solvers not to publish them.

#!/usr/bin/env bash
# Download your personal puzzle input for a given day into <year>/dayNN/input.txt.
#
# Usage: scripts/get.sh <day> [year]
#   scripts/get.sh 1        -> 2025/day01/input.txt
#   scripts/get.sh 7 2024   -> 2024/day07/input.txt
#
# Requires your Advent of Code session cookie in a file named `.session` at the
# repo root (gitignored). Get it from: log in at adventofcode.com, open browser
# DevTools -> Application/Storage -> Cookies -> copy the value of `session`.
set -euo pipefail

if [[ $# -lt 1 ]]; then
	echo "usage: $0 <day> [year]" >&2
	exit 1
fi

day10=$((10#$1))                    # strip any leading zero, force base 10
day=$(printf "%02d" "$day10")
year="${2:-2025}"
root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
session_file="$root/.session"
out="$root/$year/day$day/input.txt"

if [[ ! -f "$session_file" ]]; then
	echo "error: no session cookie found at $session_file" >&2
	echo "  create it with your AoC 'session' cookie value:" >&2
	echo "    echo 'YOUR_SESSION_COOKIE' > $session_file" >&2
	exit 1
fi

session="$(tr -d '[:space:]' < "$session_file")"
if [[ -z "$session" ]]; then
	echo "error: $session_file is empty" >&2
	exit 1
fi

mkdir -p "$(dirname "$out")"

url="https://adventofcode.com/$year/day/$day10/input"
# AoC etiquette: identify yourself in the User-Agent.
ua="github.com/utkarshraii/advent by utkarshr348@gmail.com"

http_code=$(curl -sS -w "%{http_code}" -o "$out.tmp" \
	-H "Cookie: session=$session" \
	-A "$ua" \
	"$url")

if [[ "$http_code" != "200" ]]; then
	echo "error: download failed (HTTP $http_code) for $url" >&2
	echo "  --- server response ---" >&2
	sed 's/^/  /' "$out.tmp" >&2 || true
	rm -f "$out.tmp"
	if [[ "$http_code" == "400" || "$http_code" == "500" ]]; then
		echo "  (a 400/500 usually means your session cookie is stale — refresh it in .session)" >&2
	elif [[ "$http_code" == "404" ]]; then
		echo "  (a 404 usually means this puzzle isn't unlocked yet)" >&2
	fi
	exit 1
fi

mv "$out.tmp" "$out"
lines=$(wc -l < "$out" | tr -d ' ')
bytes=$(wc -c < "$out" | tr -d ' ')
echo "saved $out ($lines lines, $bytes bytes)"

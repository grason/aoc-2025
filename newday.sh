#!/usr/bin/env bash
set -euo pipefail

# ---- Get day number -------------------------------------------------
if [[ $# -eq 1 ]] && [[ $1 =~ ^[0-9]{1,2}$ ]]; then
  DAY=$1
elif [[ $(date +%m) == "12" ]]; then
  DAY=$(date +%-d)
else
  read -p "Day number (1–25): " DAY
fi

[[ $DAY -ge 1 && $DAY -le 25 ]] || { echo "Invalid day"; exit 1; }

DIR=$(printf "Day_%02d" "$DAY")

[[ ! -e "$DIR" ]] || { echo "$DIR already exists!"; exit 1; }

echo "Creating $DIR ..."

mkdir "$DIR"
cd "$DIR"

# Empty input files
touch input.txt
touch test_input.txt

# Four completely blank solution files (you pick the extension later)
echo "package main" > solution.go
cat >> solution.go << 'EOF'
import (
	"fmt"
	"os"
	"strings"
)
func Part1(input string) int {

	return 0
}
func Part2(input string) int {

	return 0
}

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", filename, err)
		os.Exit(1)
	}

	input := string(data)
	//end boilerplate

	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}
EOF




echo "package main" > solution_test.go
cat >> solution_test.go << 'EOF'
import (
	"testing"
)
func TestPart1(t *testing.T) {
	input := ``
	got := Part1(input)
	want := 3
	if got != want {
		t.Errorf("Part1 = %d; want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := ''
	got := Part2(input)
	want := 6
	if got != want {
		t.Errorf("Part2 = %d; want %d", got, want)
	}
}
EOF


# The most generic Makefile ever written
cat > Makefile << 'EOF'
.PHONY: all run-part1 run-part2 test

all: test run\

run: run-go

test: test-go

test-go:
	go test

run-go:
	go run solution.go input.txt
EOF

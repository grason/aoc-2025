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
		tests := []struct {
		name string
		input string
		want int
	}{
		{"1","2",3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part1(tt.input)
			if got != tt.want {
				t.Errorf("Add(%s) = %d; want %d", tt.input, got, tt.want)
			}
		})
	}
}
func TestPart2(t *testing.T) {
		tests := []struct {
		name string
		input string
		want int
	}{
		{"1","2",3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part2(tt.input)
			if got != tt.want {
				t.Errorf("Add(%s) = %d; want %d", tt.input, got, tt.want)
			}
		})
	}
}
EOF


# The most generic Makefile ever written
cat > Makefile << 'EOF'
.PHONY: all test run build benchmark

all: test run
test:
	go test
run-go:
	go run solution.go input.txt
build:
	go build solution.go
benchmark: build
	time -p ./solution
EOF

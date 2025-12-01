package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

// Proper mathematical modulo — always returns 0 ≤ result < m
func mod(a, m int) int {
	return (a % m + m) % m
}
// helper so `go run . test` works (optional but nice)
func matchString(pat, str string) (bool, error) { return pat == str, nil }

func TestPart1(t *testing.T) {
	filename := "test_input.txt"
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", filename, err)
		os.Exit(1)
	}

	input := string(data)
	got := Part1(input)
	want := 3
	if got != want {
		t.Fatalf("Part1() = %d; want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	filename := "test_input.txt"
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", filename, err)
		os.Exit(1)
	}

	input := string(data)

	got := Part2(input)
	want := 6
	if got != want {
		t.Fatalf("Part2() = %d; want %d", got, want)
	}
}


// Part1 counts how many times the dial lands exactly on 0 after a full rotation
func Part1(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	dial := 50
	zeros := 0

	for _, line := range lines {
		if line == "" {
			continue
		}
		dir := line[0]
		var steps int
		fmt.Sscanf(line[1:], "%d", &steps)

		if dir == 'L' {
			dial = mod(dial-steps, 100)
		} else {
			dial = mod(dial+steps, 100)
		}
		if dial == 0 {
			zeros++
		}
	}
	return zeros
}

// Part2 counts every single time the dial passes through 0, click by click
func Part2(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	dial := 50
	clicks := 0

	for _, line := range lines {
		if line == "" {
			continue
		}
		dir := line[0]
		var steps int
		fmt.Sscanf(line[1:], "%d", &steps)

		step := 1
		if dir == 'L' {
			step = -1
		}
		for i := 0; i < steps; i++ {
			dial = mod(dial+step, 100)
			if dial == 0 {
				clicks++
			}
		}
	}
	return clicks
}

func test() {
	testing.Main(matchString, []testing.InternalTest{
		{"TestPart1", TestPart1},
		{"TestPart2", TestPart2},
	}, nil, nil)
	return
}

func main() {
	if os.Args[1] == "test" {
			test()
			return
		}


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

	fmt.Println("Part 1:", Part1(input)) // → 1177
	fmt.Println("Part 2:", Part2(input)) // → 6768
}
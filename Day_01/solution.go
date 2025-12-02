// solution.go
package main

import (
	"fmt"
	"os"
	"strings"
)

// Proper mathematical modulo — always ≥ 0
func mod(a, m int) int {
	return (a%m + m) % m
}

// Part 1: count how many times we land exactly on 0 after a full rotation
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

// Part 2: count every single click that passes through 0
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

// ————————————————————————
// main() — only runs the program
// ————————————————————————

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
	fmt.Println("Part 1:", Part1(input)) // → 1177
	fmt.Println("Part 2:", Part2(input)) // → 6768
}

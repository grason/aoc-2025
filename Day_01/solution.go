package main

import (
	"fmt"
	"os"
	"strings"
)
//go's % is wrong?
func mod(a, m int) int {
    return (a % m + m) % m
}

func main() {
	// Allow running with: go run . input.txt  OR  go run solution.go input.txt
	filename := "input.txt"
	if len(os.Args) >= 2 {
		filename = os.Args[1]
	}

	data, _ := os.ReadFile(filename)
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	dial1 := 50
	zeros := 0

	dial2 := 50
	clicks := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		dir := line[0] // 'L' or 'R'
		var steps int
		fmt.Sscanf(line[1:], "%d", &steps)

		// === Part 1: big jumps (only count when we land on 0) ===
		if dir == 'L' {
			dial1 = mod(dial1-steps, 100)
		} else {
			dial1 = mod(dial1+steps, 100)
		}
		if dial1 == 0 {
			zeros++
		}

		// === Part 2: every single click (count every time we pass through 0) ===
		step := 1
		if dir == 'L' {
			step = -1
		}
		for i := 0; i < steps; i++ {
			dial2 = mod(dial2+step, 100)
			if dial2 == 0 {
				clicks++
			}
		}
	}

	fmt.Println("Part 1:", zeros)  // → 1177
	fmt.Println("Part 2:", clicks) // → 6768
}
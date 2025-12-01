package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: go run solution.go <input.txt>")
		os.Exit(1)
	}
	filename := os.Args[1]

	file, _ := os.ReadFile(filename)
	lines := strings.Split(strings.TrimSpace(string(file)), "\n")
	dial1, zeros := 50, 0
	dial2, clicks := 50, 0

	for _, line := range lines {
		if line == "" { continue }
		dir, n := line[0], 0
		fmt.Sscanf(line[1:], "%d", &n)

		// Part 1 – big jumps
		if dir == 'L' {
			dial1 = (dial1 - n%100 + 10000) % 100
		} else {
			dial1 = (dial1 + n) % 100
		}
		if dial1 == 0 { zeros++ }

		// Part 2 – every single click
		step := 1
		if dir == 'L' { step = -1 }
		for i := 0; i < n; i++ {
			dial2 = (dial2 + step + 100) % 100
			if dial2 == 0 { clicks++ }
		}
	}

	fmt.Println("Part 1:", zeros)   // 1177
	fmt.Println("Part 2:", clicks)  // 6768
}
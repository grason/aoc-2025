package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func splitAndTrim(s, sep string) []string {
	parts := strings.Split(s, sep)
	var result []string
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" { // skip empty parts
			result = append(result, trimmed)
		}
	}
	return result
}

func StringTo2DChars(input string) [][]rune {
	// Split the input into lines
	lines := strings.Split(strings.TrimSuffix(input, "\n"), "\n")

	// Convert each line to []rune
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	return grid
}

func parseInput(input string) ([][]int, []string, int, int) {
	lines := strings.Split(input, "\n")
	l1split := strings.Split(lines[0], " ")
	width := len(l1split)
	height := len(lines) - 1
	operators := []string{}
	grid := make([][]int, height) // create slice of rows
	for i := range grid {
		grid[i] = make([]int, width) // each row has 'width' columns, all zero
	}
	//fmt.Println(width, height, len(grid[0]), len(grid))

	for j, line := range lines {
		linesplit := splitAndTrim(line, " ")
		for k, item := range linesplit {
			if j == height {
				return grid, linesplit, width, height
			}
			//fmt.Println(j, k)
			grid[j][k], _ = strconv.Atoi(item)
		}
	}

	return grid, operators, width, height
}

func Part1(input string) int {
	grid, ops, width, height := parseInput(input)
	//fmt.Println(grid, ops)

	total := 0
	for w := 0; w < width && w < len(ops); w++ {
		//fmt.Println(len(ops))
		add := ops[w] == "+"
		sum := 0
		for h := 0; h < height; h++ {
			if add {
				sum += grid[h][w]
			} else {
				if h == 0 {
					sum = grid[h][w]
				} else {
					sum *= grid[h][w]
				}
			}
		}
		total += sum
	}

	return total
}

func Part2(input string) int {
	runes := StringTo2DChars(input)
	fmt.Println(runes)

	for x := len(runes[0]) - 1; x >= 0; x-- {
		ru := make([]rune, len(runes)-1)
		for y := 0; y < len(runes)-1; y++ {
			ru[y] = runes[y][x]
		}
		fmt.Println(string(ru))
	}
	return 0
}

func main() {
	filename := "input.txt"
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	input := strings.TrimSpace(string(data))

	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}

package main

import (
	"fmt"
	"os"
	"strings"
)

func to2darray(input string) [][]int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	rows := len(lines)
	cols := len(lines[0])

	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	for y, line := range lines {
		for x, r := range line {
			//fmt.Println(y, x)
			if r == '@' {
				matrix[y][x] = 1
			} else {
				matrix[y][x] = 0
			}
		}
	}

	return matrix
}

func checkAccessable(grid [][]int, col int, row int) bool {
	sum := 0
	for x := (col - 1); x <= (col + 1); x++ {
		for y := (row - 1); y <= (row + 1); y++ {
			if x < 0 || y < 0 || y >= len(grid) || x >= len(grid[y]) {
				continue
			}
			//fmt.Println("@", y, x, "grid->", grid[y][x])
			sum += grid[y][x]
		}
	}
	//fmt.Println("@", row, col, "->", sum)
	return sum < 5 //we are counting ourselves too.
}

func Part1(input string) int {
	grid := to2darray(input)
	//fmt.Println(grid)
	sum := 0
	for y, row := range grid {
		for x, _ := range row {
			if grid[y][x] != 0 && checkAccessable(grid, x, y) {
				sum++
			}
		}
	}

	return sum
}

func Part2(input string) int {
	grid := to2darray(input)
	//fmt.Println(grid)
	sum := 0
	exit := false
	for !exit {
		exit = true
		for y, row := range grid {
			for x, _ := range row {
				if grid[y][x] != 0 && checkAccessable(grid, x, y) {
					grid[y][x] = 0
					exit = false
					sum++
				}
			}
		}
	}

	return sum
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

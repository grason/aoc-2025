package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseInput(input string) ([][2]int, []int) {

	parts := strings.Split(input, "\n\n")
	half1 := strings.Split(parts[0], "\n")
	half2 := strings.Split(parts[1], "\n")
	pairs := make([][2]int, len(half1))
	items := make([]int, len(half2))

	for y, line := range half1 {
		vals := strings.Split(line, "-")
		pairs[y][0], _ = strconv.Atoi(vals[0])
		pairs[y][1], _ = strconv.Atoi(vals[1])
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][0] != pairs[j][0] {
			return pairs[i][0] < pairs[j][0] // primary sort
		}
		return pairs[i][1] < pairs[j][1] // tie breaker
	})

	for y, line := range half2 {
		items[y], _ = strconv.Atoi(line)
	}

	return pairs, items
}

func Part1(input string) int {

	pairs, items := parseInput(input)
	fresh := 0
	for _, val := range items {
		for p := 0; p < len(pairs); p++ {
			//fmt.Println(i, pairs[p][0], val, pairs[p][1])
			if pairs[p][0] <= val && val <= pairs[p][1] {
				fresh++
				break
			}
		}
	}
	return fresh
}
func countCoveredNumbers(ranges [][2]int) int {
	if len(ranges) == 0 {
		return 0
	}

	// Create events: +1 at start, -1 at end+1
	events := make([][2]int, 0, len(ranges)*2)
	for _, r := range ranges {
		start := r[0]
		end := r[1] + 1                           // exclusive end for the sweep
		events = append(events, [2]int{start, 1}) // start event
		events = append(events, [2]int{end, -1})  // end event
	}

	// Sort events: primarily by value, secondarily process starts (+1) before ends (-1) if at same point
	sort.Slice(events, func(i, j int) bool {
		if events[i][0] != events[j][0] {
			return events[i][0] < events[j][0]
		}
		return events[i][1] < events[j][1] // +1 before -1
	})

	count := 0
	active := 0 // number of overlapping ranges currently
	prev := 0   // previous event position

	for _, e := range events {
		val, typ := e[0], e[1]

		// If there was coverage before this event, add the length of the segment
		if active > 0 {
			count += val - prev
		}

		active += typ

		// Update prev to current position
		// (for the first event, this segment addition is skipped since active was 0)
		prev = val
	}

	return count
}
func Part2(input string) int {
	pairs, _ := parseInput(input)
	return countCoveredNumbers(pairs)
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

package main

import "testing"

func TestPart1(t *testing.T) {
	input := "..@@.@@@@.\n" +
		"@@@.@.@.@@\n" +
		"@@@@@.@.@@\n" +
		"@.@@@@..@.\n" +
		"@@.@@@@.@@\n" +
		".@@@@@@@.@\n" +
		".@.@.@.@@@\n" +
		"@.@@@.@@@@\n" +
		".@@@@@@@@.\n" +
		"@.@.@@@.@.\n"
	want := 13
	got := Part1(input)
	if got != want {
		t.Errorf("Part1() = %d; want 13", got)
	}
}

func TestPart2(t *testing.T) {
	// Add tests here
}

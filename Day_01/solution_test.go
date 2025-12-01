package day01

import (
	"testing"
)
func TestPart1(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

	got := Part1(input)
	want := 3
	if got != want {
		t.Errorf("Part1 = %d; want %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

	got := Part2(input)
	want := 6
	if got != want {
		t.Errorf("Part2 = %d; want %d", got, want)
	}
}
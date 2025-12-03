// solution_test.go
package main

import (
	"reflect"
	"testing"
)

func TestPart1_ExampleInput(t *testing.T) {
	input := `
987654321111111
811111111111119
234234234234278
818181911112111
`

	expected := 98 + 89 + 78 + 92 // 357
	got := Part1(input)

	if got != expected {
		t.Errorf("Part1() = %d; want %d", got, expected)
	}
}

func TestPart1_SingleRow(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "only one battery",
			input:    "5",
			expected: 0, // can't form two-digit number
		},
		{
			name:     "two batteries: 98",
			input:    "98",
			expected: 98,
		},
		{
			name:     "two batteries: 89",
			input:    "89",
			expected: 89,
		},
		{
			name:     "increasing order",
			input:    "123456789",
			expected: 89,
		},
		{
			name:     "9 at start, 8 at end",
			input:    "911111118",
			expected: 98,
		},
		{
			name:     "8 at start, 9 at end",
			input:    "8111111119",
			expected: 89,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(tt.input); got != tt.expected {
				t.Errorf("Part1() = %d; want %d", got, tt.expected)
			}
		})
	}
}

func Test_maxJolts(t *testing.T) {
	tests := []struct {
		name string
		bank []int
		want int
	}{
		{
			name: "987654321111111 → 98",
			bank: []int{9,8,7,6,5,4,3,2,1,1,1,1,1,1,1},
			want: 98,
		},
		{
			name: "811111111111119 → 89",
			bank: []int{8,1,1,1,1,1,1,1,1,1,1,1,1,1,9},
			want: 89,
		},
		{
			name: "234234234234278 → 78",
			bank: []int{2,3,4,2,3,4,2,3,4,2,3,4,2,7,8},
			want: 78,
		},
		{
			name: "818181911112111 → 92",
			bank: []int{8,1,8,1,8,1,9,1,1,1,1,2,1,1,1},
			want: 92,
		},
		{
			name: "Just 1s → 11",
			bank: []int{1,1,1,1},
			want: 11,
		},
		{
			name: "single 9",
			bank: []int{9},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lookup := generateLookup(tt.bank)
			if got := maxJolts(lookup); got != tt.want {
				t.Errorf("maxJolts() = %d; want %d", got, tt.want)
			}
		})
	}
}

func Test_generateLookup(t *testing.T) {
	bank := []int{9,8,7,6,5,4,3,2,1,1,1,1,1,1,1}
	lookup := generateLookup(bank)

	// Check that position 0 has 9, position 1 has 8, etc.
	if len(lookup[8]) == 0 || lookup[8][0] != 0 { // '9' → index 8
		t.Errorf("Expected battery 9 at position 0")
	}
	if len(lookup[7]) == 0 || lookup[7][0] != 1 { // '8' → index 7
		t.Errorf("Expected battery 8 at position 1")
	}
}

func Test_assembleBank(t *testing.T) {
	input := "987\n654\n321"
	ch := make(chan []int, 3)

	go assembleBank(input, ch)

	// Collect all rows
	var results [][]int
	for row := range ch {
		results = append(results, row)
	}

	expected := [][]int{
		{9, 8, 7},
		{6, 5, 4},
		{3, 2, 1},
	}

	if len(results) != 3 {
		t.Fatalf("Expected 3 rows, got %d", len(results))
	}

	for i, exp := range expected {
		if !reflect.DeepEqual(results[i], exp) {
			t.Errorf("Row %d: got %v, want %v", i, results[i], exp)
		}
	}
}

func Test_Part1_EmptyInput(t *testing.T) {
	if got := Part1(""); got != 0 {
		t.Errorf("Part1(empty) = %d; want 0", got)
	}
}

func Test_Part1_OneEmptyLine(t *testing.T) {
	if got := Part1("\n"); got != 0 {
		t.Errorf("Part1(one empty line) = %d; want 0", got)
	}
}
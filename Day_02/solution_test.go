package main

import (
	"testing"
	"slices"
)

func TestSplitMatch(t *testing.T) {
		tests := []struct {
		name string
		input int
		want bool
	}{
		{"small",1,false},
		{"large", 123123123123123, false},
		{"match", 123123123123, true},
		{"22", 22, true},
		{"123454321", 123454321, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := splitMatch(tt.input)
			if got != tt.want {
				t.Errorf("Add(%d) = %t; want %t", tt.input, got, tt.want)
			}
		})
	}
}
func TestSearch(t *testing.T) {
		tests := []struct {
		name string
		txt string
		pat string
		want []int
	}{
		{"TestSearchExample","aabaacaadaabaaba","aaba",[]int{0, 9, 12}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := search(tt.pat, tt.txt)
			if !slices.Equal(got,tt.want) {
				t.Fail()
				//t.Errorf("Add(%s) = %d; want %d", tt.input, got, tt.want)
			}
		})
	}
}


func TestPart1(t *testing.T) {
		tests := []struct {
		name string
		input string
		want int
	}{
		{"TestPart1Example","11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124",1227775554},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part1(tt.input)
			if got != tt.want {
				t.Errorf("Add(%s) = %d; want %d", tt.input, got, tt.want)
			}
		})
	}
}
func TestPart2(t *testing.T) {
		tests := []struct {
		name string
		input string
		want int
	}{
		{"TestPart2Example","11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124",4174379265},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part2(tt.input)
			if got != tt.want {
				t.Errorf("Add(%s) = %d; want %d", tt.input, got, tt.want)
			}
		})
	}
}
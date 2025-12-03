package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func assembleBank(input string, output chan<- []int){
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		arr := make([]int,len(line),len(line))
		for pos, r := range line {
			arr[pos],_ = strconv.Atoi(string(r))
		}
		output <- arr
	}
	close(output)
}

func generateLookup(bank []int)[9][]int{
	var positions [9][]int
	for pos, battery := range bank {
		if battery == 0 {
			continue
		}
		//fmt.Println(pos, battery)
		positions[battery-1] = append(positions[battery-1], pos)
	}
	return positions
}


func maxJolts(look [9][]int) int {
	//first and lasts and valid configurations
	//the list of firsts
	var first [9]int
	var last [9]int
	for i := 0; i < 9; i++ {
		if len(look[i]) > 0{
			first[i] = look[i][0]+1
			last[i] = look[i][len(look[i])-1]+1
		}
	}
	//fmt.Println("First: ", first)
	//fmt.Println("Last: ", last)
	for i := 8; i >=0; i-- {
		if first[i] == 0 {
			continue
		}
		for j := 8; j >=0; j-- {
			if last[j] == 0 {
				continue
			}
			if (first[i] < last[j]){
				result := (i+1)*10 + (j+1)
				//fmt.Println(result)
				return result

			}
		}
	}
	return 0
}

func calculateMaxJolts(input <-chan []int, output chan<- int){
//array of slices, battery positions by power
	for bank := range input {
		//fmt.Println(bank)
		look := generateLookup(bank)
		//fmt.Println(look)
		jolts := maxJolts(look);
		output <- jolts
	}
	close(output)
}


func Part1(input string) int {

	//make out channel for assembled battery banks.
	banks_ch := make(chan []int)
	go assembleBank(input, banks_ch)
	maxjolt_ch := make(chan int)
	go calculateMaxJolts(banks_ch, maxjolt_ch)

	sum := 0
	for j := range maxjolt_ch{
		sum += j
	}
	return sum
}

func Part2(input string) int {
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


//offline elevators
//joltage
//each line is a bank
//within each bacnk we need to turn on an arbitrary number of batteries (2)
//the two digit number that the turned on batteries make is the jotable
//find the largest possible joltage each bank can produce. 


//probably a pointer for each digit in th joltage
//a data structure that looks something like 

// -> a battery canidate with an array of sup canidates?

// would it be better to make an arrangement of the first index of each number?
// first 9: index 3
// first 8: index 2
// first 7: index 4


// or a list of number locations, 9 indices: [3 10 11 12]
// normally you would just take the first. 


//then when it comes time to calculate, we look at 9, find first index. then look at 8, is the last index greater than 9's first index?
// if not move on.




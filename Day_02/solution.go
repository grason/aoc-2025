package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//splits ints in the middle, then returns if they match.
func splitMatch(input int) bool {
	inputStr := strconv.Itoa(input);
	ilen := len(inputStr)
	if ilen % 2 != 0 {
		return false
	}
	half := ilen / 2
	return inputStr[:half] == inputStr[half:]
}

// stuffs ints into a channel, then closes channel.
func rangeProducer(input string, output chan<- int){
	for _,s := range strings.Split(input, ",") {
		s_arr := strings.Split(s,"-")
		start, _ := strconv.Atoi(s_arr[0])
		end, _ := strconv.Atoi(s_arr[1])
		for i := start; i <= end; i++ {
			output <- i
		}
	}
	close(output)
}

func isRepeatingNumber(num int) bool {
    s := strconv.Itoa(num)
    n := len(s)

    for i := 1; i <= n/2; i++ {
        if n%i != 0 {
            continue
        }
        if strings.Repeat(s[:i], n/i) == s {
            return true
        }
    }
    return false
}


func verify(inputch <-chan int, invalidIds chan<- int){
	defer close(invalidIds)
	for num := range inputch {
		if splitMatch(num) {
			invalidIds <- num
		}
	}
}



func verify2(inputch <-chan int, invalidIds chan<- int){
	defer close(invalidIds)
	for num := range inputch {
		if isRepeatingNumber(num) {
			invalidIds <- num
		}
	}
}


func Part1(input string) int {

	//we doing some ~golang stuff~ now
	inputch := make(chan int, 32)
	invalidIds := make(chan int)

	go rangeProducer(input, inputch)
	go verify(inputch, invalidIds)

	sum := 0

	for n := range invalidIds {
		sum += n
	}
	return sum
	
}
func Part2(input string) int {

	//we doing some golang stuff now
	inputch := make(chan int, 32)
	invalidIds := make(chan int)

	go rangeProducer(input, inputch)
	go verify2(inputch, invalidIds)

	sum := 0

	for n := range invalidIds {
		sum += n
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
		fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", filename, err)
		os.Exit(1)
	}

	input := string(data)
	//end boilerplate

	fmt.Println("Part 1:", Part1(input))
	fmt.Println("Part 2:", Part2(input))
}

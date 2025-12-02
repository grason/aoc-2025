package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"slices"
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

//God bless Donald Knuth
//https://www.geeksforgeeks.org/dsa/kmp-algorithm-for-pattern-searching/
func lpsArray(pattern string)[]int {
	length := 0
	lps := make([]int, len(pattern))
	i := 1
	for i < len(pattern) {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

func search(pat string, txt string)[]int{
	n:= len(txt)
	m:= len(pat)
	res :=[]int{}
	lps := lpsArray(pat)
	i := 0
	j := 0
	for i < n {
		if txt[i] == pat[j] {
			i++
			j++
			if j == m {
				res = append(res, i - j)
				j = lps[j-1]
			}
		} else {
			if j != 0 {
                j = lps[j-1]
			} else {
                i++;
            }
		}
	}
	return res
}


// the reference slice will look like this ALWAYS
// Example 12121212
// ref [0 2 4 6]
func refRes(patSize int, txtSize int)[]int{
	arrSize := txtSize / patSize
	rval := make([]int, arrSize)
	for i:=0;i*patSize<txtSize;i++{
		rval[i] = i*patSize
	}
	return rval
}

//we find the kmp array for this number, then we must check the returns 
func kmpMatch(num int) bool {
	inputStr := strconv.Itoa(num);
	//we only need to check the first half of of the number for prefixes.
	for i:=1; i<=len(inputStr)/2;i++ {
		if len(inputStr) % i != 0 {
			continue //skip patterns that don't equally fit within the number
		}
		//construct lps
		res := search(inputStr[:i], inputStr)
		//construct reference slice which we compare the actual res slice to
		ref := refRes(i,len(inputStr))

		if slices.Equal(res,ref) {
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
		if kmpMatch(num) {
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

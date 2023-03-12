package main

import "fmt"

/**
Given an array of integers, find the one that appears an odd number of times.

There will always be only one integer that appears an odd number of times.

Examples
[7] should return 7, because it occurs 1 time (which is odd).
[0] should return 0, because it occurs 1 time (which is odd).
[1,1,2] should return 2, because it occurs 1 time (which is odd).
[0,1,0,1,0] should return 0, because it occurs 3 times (which is odd).
[1,2,2,3,3,3,4,3,3,3,2,2,1] should return 4, because it appears 1 time (which is odd).
*/

func FindOdd(seq []int) int {
	letters := make(map[string]interface{})
	if len(seq) == 1 {
		return seq[0]
	}

	for l := 0; l < len(seq); l++ {
		fmt.Println(seq[l])
		if val, ok := letters[seq[l]]; ok {
			fmt.Println("Hello", val)
		}
	}

	return 1000
}

func main() {
	result := FindOdd([]int{7, 3})
	fmt.Println(result)
}

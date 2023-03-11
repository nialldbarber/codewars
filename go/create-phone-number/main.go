package main

import (
	"fmt"
	"strconv"
)

/**
	Write a function that accepts an array
	of 10 integers (between 0 and 9), that
	returns a string of those numbers in the
	form of a phone number

	CreatePhoneNumber([10]uint{1,2,3,4,5,6,7,8,9,0})
	// returns "(123) 456-7890"
**/

func convertNumbersToString(numbers []uint) string {
	var result string
	for _, number := range numbers {
		result += strconv.Itoa(int(number))
	}
	return result
}

func CreatePhoneNumber(numbers [10]uint) string {
	firstChunk := convertNumbersToString(numbers[0:3])
	secondChunk := convertNumbersToString(numbers[3:6])
	thirdChunk := convertNumbersToString(numbers[6:10])
	result := fmt.Sprintf("(%s) %s-%s", firstChunk, secondChunk, thirdChunk)
	return result
}

func main() {
	var result = CreatePhoneNumber([10]uint{1,2,3,4,5,6,7,8,9,0})
	fmt.Println(result)
}

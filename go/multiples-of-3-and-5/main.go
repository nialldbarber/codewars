package main

import "fmt"

func Multiple3And5(number int) int {
	total := 0
	multiples := []int{}

	for i := 0; i < number; i++ {
		if i % 3 == 0 {
			multiples = append(multiples, i)
		} else if i % 5 == 0 {
			multiples = append(multiples, i)
		}
	}

	for i := 0; i < len(multiples); i++ {
		total += multiples[i]
	}

	return total
}

func main() {
	result := Multiple3And5(10)
	fmt.Println(result)
}

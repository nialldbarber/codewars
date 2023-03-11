package main

import "fmt"

/*
   Create a function that takes an
   integer as an argument and returns
   "Even" for even numbers or "Odd"
   for odd numbers.
*/


func EvenOrOdd(number int) string {
  if number % 2 == 0 {
		return "Even"
	}
	return "Odd"
}

func main() {
	fmt.Println(EvenOrOdd(10))
	fmt.Println(EvenOrOdd(5))
}

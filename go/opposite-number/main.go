package main

import "fmt"

func Opposite(value int) int {
  return value * -1
}

func main() {
	response := Opposite(10)
	fmt.Println(response)
}

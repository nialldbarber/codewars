package main

import (
	"fmt"
	"strings"
)

func GetCount(str string) (count int) {
	vowels := []string{"a", "e", "i", "o", "u"}
	splitStrings := strings.Split(str, "")
	
	for _, s1 := range splitStrings {
		for _, s2 := range vowels {
			if s1 == s2 {
				count += 1
			}
		}
	}

  return count
}

func main() {
	fmt.Println(GetCount("hello"))
}

package main

import (
	"fmt"
	"strings"
)

func Disemvowel(comment string) string {
	outcome := []string{}
	vowels := []string{"a", "e", "i", "o", "u"}
	splitStrings := strings.Split(comment, "")
	for _, s1 := range splitStrings {
		for _, s2 := range vowels {
			if s2 == s1 {
				outcome = append(outcome, s1)
			}
		}
	}
	return strings.Join(outcome, "")
}

func main() {
	fmt.Println(Disemvowel("hello"))
}

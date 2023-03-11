package main

import (
	"fmt"
	"strings"
)

func Disemvowel(comment string) string {
	vowels := []string{"a", "A", "e", "E", "i", "I", "o", "O", "u", "U"}
	for _, letter := range vowels {		
		comment = strings.ReplaceAll(comment, string(letter), "")
	}
	return comment
}

func main() {
	fmt.Println(Disemvowel("hello"))
	fmt.Println(Disemvowel("this is a test LOL, HAppy daYs"))
}

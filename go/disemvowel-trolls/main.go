package main

import (
	"fmt"
	"strings"
)

func Disemvowel(comment string) string {
	for _, letter := range vowels.Vowels {		
		comment = strings.ReplaceAll(comment, string(letter), "")
	}
	return comment
}

func main() {
	fmt.Println(Disemvowel("hello"))
	fmt.Println(Disemvowel("this is a test LOL, HAppy daYs"))
}

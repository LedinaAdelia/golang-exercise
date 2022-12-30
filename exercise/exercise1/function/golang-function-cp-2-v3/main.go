package main

import (
	"fmt"
	"strings"
)

func CountVowelConsonant(str string) (int, int, bool) {
	vowel := 0
	consonant := 0

	splitChar := strings.Split(str, "")
	for _, v := range splitChar {
		checkChar := strings.ToUpper(v)
		if strings.Contains("AIUEO", checkChar) {
			vowel++
		} else if !strings.Contains(" ,", checkChar) {
			consonant++
		}
	}

	if vowel == 0 {
		return vowel, consonant, true
	}
	return vowel, consonant, false
}

func main() {
	fmt.Println(CountVowelConsonant("bbbbb, ccccc"))
}

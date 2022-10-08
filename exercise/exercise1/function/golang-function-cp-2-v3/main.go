package main

import (
	"fmt"
	"strings"
)

func CountVowelConsonant(str string) (int, int, bool) {
	arrString := strings.Split(str, "")
	countVok := 0
	countKon := 0
	konEmpty := false
	for i := 0; i < len(arrString); i++ {
		char := strings.ToUpper(arrString[i])
		if char == "A" || char == "I" || char == "U" || char == "E" || char == "O" {
			countVok++
		} else if char == " " {
			continue
		} else {
			countKon++
		}
	}
	if countVok == 0 {
		konEmpty = true
	} else {
		konEmpty = false
	}

	if strings.Contains(str, ",") {
		countKon--
	}

	return countVok, countKon, konEmpty // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountVowelConsonant("Hidup Itu Indah"))
	fmt.Println(CountVowelConsonant("SEMANGAT PAGI, itu kata orang yang baru datang ke rumahku"))
}

package main

import (
	"fmt"
	"strings"
)

func SlurredTalk(words *string) {
	arrGenerate := strings.Split(*words, "")
	vokal := []string{"s", "r", "z", "S", "R", "Z"}
	output := ""
	for i := 0; i < len(arrGenerate); i++ {
		a := arrGenerate[i]
		for j := 0; j < len(vokal); j++ {
			if a == vokal[j] {
				if strings.ToUpper(a) == a {
					a = "L"
				} else {
					a = "l"
				}

				break
			}
		}

		output += a
	}
	*words = output
}

func main() {
	var words string = "Streven is meZ"
	SlurredTalk(&words)
	fmt.Println(words)
}

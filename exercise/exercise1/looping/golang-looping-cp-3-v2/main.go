package main

import (
	"fmt"
	"strings"
)

func CountingLetter(text string) int {
	count := 0
	for i := 0; i < len(text); i++ {
		char := strings.ToUpper(string(text[i]))
		if strings.Contains("RSTZ", char) {
			count++
		}
	}
	return count

}

func main() {
	fmt.Println(CountingLetter("Semangat"))
}

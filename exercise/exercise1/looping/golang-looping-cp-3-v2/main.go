package main

import (
	"fmt"
	"strings"
)

func CountingLetter(text string) int {
	r := []rune(text)
	count := 0
	for i := 0; i < len(r); i++ {
		char := strings.ToUpper(string(r[i]))
		if char == "R" || char == "S" || char == "T" || char == "Z" {
			count++
		}
	}
	return count

}

func main() {
	fmt.Println(CountingLetter("Semangat"))
}

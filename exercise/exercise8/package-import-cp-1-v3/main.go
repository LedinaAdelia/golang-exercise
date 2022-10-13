package main

import (
	"fmt"
	"strconv"
	"strings"
)

func AdvanceCalculator(calculate string) float32 {
	a := strings.Split(calculate, " ")
	fmt.Println(a)
	cal := 0
	for i := 0; i < len(a); i++ {
		if a[i] == "*" {
			if cal == 0 {
				bil1, _ := strconv.Atoi(a[i-1])
				bil2, _ := strconv.Atoi(a[i+1])
				cal = bil1 * bil2
				
			} else {

			}
		} else if a[i] == "/" {

		}
	}
	fmt.Println(cal)
	return 0 // TODO: replace this
}

func main() {
	res := AdvanceCalculator("3 * 4 / 2 + 10 - 5")

	fmt.Println(res)
}

package main

import (
	"fmt"
)

func ExchangeCoin(amount int) []int {
	hm := []int{}
	// arrPiece := []int{1000, 500, 200, 100, 50, 20, 10, 5}
	count := false
	for (amount >= 0) && !count {
		if amount >= 1000 {
			amount -= 1000
			hm = append(hm, 1000)
		} else if (amount >= 500) && (amount < 1000) {
			amount -= 500
			hm = append(hm, 500)
		} else if (amount >= 200) && (amount < 500) && (amount < 1000) {
			amount -= 200
			hm = append(hm, 200)
		} else if (amount >= 100) && (amount < 200) && (amount < 500) && (amount < 1000) {
			amount -= 100
			hm = append(hm, 100)
		} else if (amount >= 50) && (amount < 100) && (amount < 200) && (amount < 500) && (amount < 1000) {
			amount -= 50
			hm = append(hm, 50)
		} else if (amount >= 20) && (amount < 50) && (amount < 100) && (amount < 200) && (amount < 500) && (amount < 1000) {
			amount -= 20
			hm = append(hm, 20)
		} else if (amount >= 10) && (amount < 20) && (amount < 50) && (amount < 100) && (amount < 200) && (amount < 500) && (amount < 1000) {
			amount -= 10
			hm = append(hm, 10)
		} else if (amount >= 5) && (amount < 10) && (amount < 20) && (amount < 50) && (amount < 100) && (amount < 200) && (amount < 500) && (amount < 1000) {
			amount -= 5
			hm = append(hm, 5)
		} else if (amount >= 1) && (amount < 5) && (amount < 10) && (amount < 20) && (amount < 50) && (amount < 100) && (amount < 200) && (amount < 500) && (amount < 1000) {
			amount -= 1
			hm = append(hm, 1)
		} else {
			count = true
		}
	}
	return hm // TODO: replace this
}
func main() {
	fmt.Println(ExchangeCoin(1752))
	fmt.Println(ExchangeCoin(5000))
	fmt.Println(ExchangeCoin(1234))
	a := ExchangeCoin(0)
	fmt.Println(a, len(a), cap(a))
}

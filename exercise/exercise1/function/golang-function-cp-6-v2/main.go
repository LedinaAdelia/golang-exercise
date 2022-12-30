package main

import (
	"fmt"
	"strconv"
)

func ChangeToCurrency(change int) string {
	arrChange := strconv.Itoa(change)
	cResultReverse := ""
	cResult := ""
	count := 0
	for i := 0; i < len(arrChange); i++ {
		if count == 2 && i != (len(arrChange)-1) {
			cResultReverse += string(arrChange[len(arrChange)-i-1]) + "."
			count = 0
		} else {
			cResultReverse += string(arrChange[len(arrChange)-i-1])
			count++
		}
	}
	for j := 0; j < len(cResultReverse); j++ {
		cResult += string(cResultReverse[len(cResultReverse)-j-1])
	}
	return fmt.Sprintf("Rp. %s", cResult)
}

func MoneyChange(money int, listPrice ...int) string {
	totalPrice := 0
	for _, v := range listPrice {
		totalPrice += v
	}
	if money < totalPrice {
		return "Uang tidak cukup"
	}
	return ChangeToCurrency(money - totalPrice)
}

func main() {
	fmt.Println(MoneyChange(100000, 50000, 10000, 10000, 5000, 5000))
}

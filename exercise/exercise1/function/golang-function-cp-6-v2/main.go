package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ChangeToCurrency(change int) string {
	arrChange := strconv.Itoa(change)
	splitChange := strings.Split(arrChange, "")
	cResultReverse := ""
	cResult := ""
	count := 0
	for i := 0; i < len(splitChange); i++ {
		if count == 2 && i != (len(splitChange)-1) {
			count = 0
			cResultReverse += splitChange[len(splitChange)-i-1] + "."

		} else {
			count++
			cResultReverse += splitChange[len(splitChange)-i-1]
		}
	}
	result := strings.Split(cResultReverse, "")
	for j := 0; j < len(result); j++ {
		cResult += result[len(result)-j-1]
	}

	return fmt.Sprintf("Rp. %s", cResult)
}

func MoneyChange(money int, listPrice ...int) string {
	arrData := []int(listPrice)
	totalPrice := 0
	result := ""
	for i := 0; i < len(arrData); i++ {
		totalPrice += arrData[i]
	}
	cData := money - totalPrice
	if cData == 0 {
		result = ChangeToCurrency(cData)
	} else if cData < 0 {
		result = "Uang tidak cukup"
	} else {
		result = ChangeToCurrency(cData)
	}
	return result // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(MoneyChange(100000, 50000, 10000, 10000, 5000, 5000))
	fmt.Println(MoneyChange(10000, 5000, 5000))
}

package main

import (
	"fmt"
	"strconv"
)

func BiggestPairNumber(numbers int) int {
	var maxValue int
	var maxValuePair string
	arrNum := strconv.Itoa(numbers)
	for i := 0; i < len(arrNum)-1; i++ {
		firstNumber, _ := strconv.Atoi(string(arrNum[i]))
		secondNumber, _ := strconv.Atoi(string(arrNum[i+1]))
		count := firstNumber + secondNumber
		if count > maxValue {
			maxValue = count
			maxValuePair = fmt.Sprintf("%d%d", firstNumber, secondNumber)
		}
	}
	pair, _ := strconv.Atoi(maxValuePair)
	return pair
}

func main() {
	fmt.Println(BiggestPairNumber(11283344))
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func BiggestPairNumber(numbers int) int {
	arrNum := strconv.Itoa(numbers)
	arrWord := strings.Split(arrNum, "")
	maxValue := 0
	maxValuePair := ""
	for i := 0; i < len(arrNum)-1; i++ {
		firstNumber, _ := strconv.Atoi(arrWord[i])
		secondNumber, _ := strconv.Atoi(arrWord[i+1])
		count := firstNumber + secondNumber
		if count > maxValue {
			maxValue = count
			maxValuePair = fmt.Sprintf("%d%d", firstNumber, secondNumber)
			// maxValueIndex = i
		}
	}
	pair, _ := strconv.Atoi(maxValuePair)
	return pair // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BiggestPairNumber(11283344))

}

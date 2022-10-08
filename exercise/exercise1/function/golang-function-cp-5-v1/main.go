package main

import (
	"fmt"
)

func FindMin(nums ...int) int {
	arrData := []int(nums)
	minNumber := arrData[0]
	minResult := arrData[0]
	for i := 1; i < len(arrData)-1; i++ {
		minData := arrData[i]
		if minData < minNumber {
			minNumber = minData
			minResult = arrData[i]
		}
	}
	return minResult // TODO: replace this
}

func FindMax(nums ...int) int {
	arrData := []int(nums)
	maxNumber := arrData[0]
	maxResult := arrData[0]
	for i := 1; i < len(arrData); i++ {
		minData := arrData[i]
		if minData > maxNumber {
			maxNumber = minData
			maxResult = arrData[i]
		}
	}
	return maxResult
}

func SumMinMax(nums ...int) int {
	arrData := []int(nums)
	min := FindMin(arrData...)
	max := FindMax(arrData...)
	result := min + max
	return result // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(SumMinMax(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

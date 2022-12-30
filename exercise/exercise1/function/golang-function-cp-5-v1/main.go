package main

import (
	"fmt"
	"sort"
)

func FindMin(nums ...int) int {
	sort.Ints(nums)
	return nums[0]
}

func FindMax(nums ...int) int {
	sort.Ints(nums)
	return nums[len(nums)-1]
}

func SumMinMax(nums ...int) int {
	min := FindMin(nums...)
	max := FindMax(nums...)
	return min + max
}

func main() {
	fmt.Println(SumMinMax(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

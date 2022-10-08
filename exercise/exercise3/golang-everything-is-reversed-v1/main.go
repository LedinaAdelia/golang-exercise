package main

import (
	"fmt"
)

func ReverseData(arr [5]int) [5]int {
	final := [5]int{}
	for i := 0; i < len(arr); i++ {
		final[i] += arr[len(arr)-i-1]
		a := 0
		reversDigits(final[i], &a)
	}
	return final // TODO: replace this
}

func reversDigits(num int, he *int) {
	rev_num := 0
	for i := 0; i < num; i++ {
		if num == 10 {
			rev_num = num / 10
		} else {
			rev_num = rev_num*10 + num%10
			num = num / 10
		}

	}
	*he = rev_num
}

func main() {
	data := [5]int{10, 10, 10, 10, 10}

	result := ReverseData(data)
	fmt.Println(result)

	data2 := [5]int{123, 456, 11, 1, 2}

	results := ReverseData(data2)
	fmt.Println(results)

}

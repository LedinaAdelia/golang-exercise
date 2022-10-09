package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ReverseData(arr [5]int) [5]int {
	final := [5]int{}
	for i := 0; i < len(arr); i++ {
		final[i] += arr[len(arr)-i-1]
		a := 0
		fmt.Println("ebelum: ", final[i])
		reversDigits(final[i], &a)
		final[i] = a

	}
	return final // TODO: replace this
}

func reversDigits(num int, he *int) {
	revNum := strconv.Itoa(num)
	fmt.Println("hai: ", revNum)
	splitNum := strings.Split(revNum, "")
	rest1 := ""
	for i := 0; i < len(splitNum); i++ {
		rest1 += splitNum[len(splitNum)-i-1]
	}
	output, _ := strconv.Atoi(rest1)
	fmt.Println("kebalik harusnya: ", output)
	*he = output
	// *he = revNum
}

func main() {
	data := [5]int{456789, 44332, 2221, 12, 10}

	result := ReverseData(data)
	fmt.Println(result)

	data2 := [5]int{123, 456, 11, 1, 2}

	results := ReverseData(data2)
	fmt.Println(results)

}

package main

import "fmt"

func SchedulableDays(date1 []int, date2 []int) []int {
	try := []int{}
	for i := 0; i < len(date1); i++ {
		for j := 0; j < len(date2); j++ {
			if date1[i] > 32 && date2[j] > 32 {
				fmt.Println("hai")
				try = nil
			} else if date1[i] == date2[j] {
				try = append(try, date1[i])
			}
		}
	}
	return try // TODO: replace this
}
func main() {
	a := []int{1, 2, 15, 6, 8}
	b := []int{2, 5, 7, 8, 10, 3, 26}
	c := SchedulableDays(a, b)
	fmt.Println(c)
}

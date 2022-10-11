package main

import (
	"fmt"
)

func SchedulableDays(villager [][]int) []int {
	res := []int{}
	a := []int{}
	for i := 0; i < len(villager); i++ {
		for j := 0; j < len(villager[i]); j++ {
			a = append(a, villager[i][j])
		}
	}
	fmt.Println(a)
	if len(villager) <= 1 {
		res = a
	} else {
		keys := make(map[int]bool)
		for _, entry := range a {
			if _, value := keys[entry]; !value {
				if !value {
					keys[entry] = true
				}
			} else {
				res = append(res, entry)
			}
		}
	}

	return res // TODO: replace this
}

func main() {
	a := [][]int{
		// {1, 2, 3, 4, 5, 6, 7, 11, 12, 13, 14, 15, 16, 17, 19, 20},
		// {3, 4, 5, 6, 7, 11, 12, 13, 14, 15, 16, 17, 19, 20},
		// {1, 2, 3, 4, 5, 6, 7, 11, 12, 13, 14, 15, 16, 17, 19, 20},
		// {15, 16, 17, 18, 19, 20},
		// {1, 2, 3, 4, 5, 6, 7, 11, 12, 13, 14, 15, 16, 17, 19, 20},
		// {1, 2, 3, 4, 5, 6, 7, 11, 12, 13, 14, 15, 16, 17, 19, 20},
		{1, 2, 3},
	}
	fmt.Println(SchedulableDays(a))
}

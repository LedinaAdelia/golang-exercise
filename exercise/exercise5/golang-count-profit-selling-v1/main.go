package main

import "fmt"

// func CountProfit(data [][][2]int) []int {
// 	r := []int{}
// 	if len(data) == 0 {

// 	} else if len(data) == 1 {
// 		sum := sumProfit(data)
// 		for _, v := range sum {
// 			r = append(r, v)
// 		}
// 	} else if len(data[0]) == 1 {
// 		sum := sumProfit(data)
// 		count := 0
// 		for _, v := range sum {
// 			count += v
// 		}
// 		r = append(r, count)
// 	} else {
// 		var res = map[int]int{}
// 		for _, a := range data {
// 			// sum1 := 0
// 			for i, b := range a {
// 				sum2 := 0
// 				for _, c := range b {
// 					if sum2 == 0 {
// 						sum2 += c
// 					} else {
// 						sum2 -= c
// 						res[i] += sum2
// 					}
// 				}
// 			}
// 		}
// 		for i := 0; i < len(res); i++ {
// 			r = append(r, res[i])
// 		}
// 	}

// 	return r
// }
// func sumProfit(data [][][2]int) []int {
// 	r := []int{}
// 	for _, a := range data {
// 		for _, b := range a {
// 			sum := 0
// 			for _, c := range b {
// 				if sum == 0 {
// 					sum += c
// 				} else {
// 					sum -= c
// 					r = append(r, sum)
// 				}
// 			}
// 		}
// 	}
// 	return r
// }

func CountProfit(data [][][2]int) {
	// res := []int{}
	for _, outside := range data {
		for _, middle := range outside {
			for _, inner := range middle {
				fmt.Println("middle, outside", len(middle), len(outside))
				fmt.Println(inner)
			}
		}
	}
}

func main() {
	// a := [][][2]int{
	// 	{
	// 		{1000, 500},
	// 		{500, 150},
	// 		{600, 100},
	// 		{800, 750},
	// 	},
	// }
	// output = 500 350 500 50

	// b := [][][2]int{
	// 	{{1000, 200}},
	// 	{{500, 100}},
	// 	{{450, 150}},
	// 	{{100, 50}},
	// }
	// output = 1550

	// c := [][][2]int{
	// 	{{1000, 800}, {700, 500}, {100, 50}},
	// 	{{1000, 800}, {900, 200}, {500, 200}},
	// 	{{1000, 900}, {900, 200}, {500, 200}},
	// }
	// output = 500 1600 650

	// CountProfit(a)
	// CountProfit(b)
	// CountProfit(c)
	// fmt.Println(CountProfit(a))
	// fmt.Println(CountProfit(b))
	// fmt.Println(CountProfit(c))
}

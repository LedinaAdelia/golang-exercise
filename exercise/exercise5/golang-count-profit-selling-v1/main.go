package main

import "fmt"

// func CountProfit(data [][][2]int) []int {
// 	a := []int{}
// 	b := []int{}
// 	c := []int{}
// 	count := 0
// 	sum := 0
// 	if len(data) == 0 {
// 		return a
// 	} else {
// 		for i := 0; i < len(data); i++ {
// 			fmt.Println("i: ", i)
// 			hum := 0
// 			for j := 0; j < len(data[i]); j++ {
// 				fmt.Println("j: ", j)
// 				sum := 0
// 				for k := 0; k < len(data[i][j]); k++ {
// 					fmt.Println("k: ", k)
// 					if sum == 0 {
// 						sum += data[i][j][k]
// 					} else {
// 						hehe := sum - data[i][j][k]
// 						count = hehe
// 						if j <= 1 {
// 							a = append(a, count)
// 						}

//						}
//					}
//					hum += count
//				}
//				b = append(b, hum)
//				sum += hum
//			}
//			c = append(c, sum)
//			fmt.Println(a)
//		}
//		return nil // TODO: replace this
//	}

func CountProfit(data [][][2]int) []int {
	a := []int{}
	sum := 0
	for i, x := range data {
		fmt.Println("i,x: ", i, x)
		for j, y := range x {
			fmt.Println("j,y: ", j, y)

			if j == 0 {
				fmt.Println("hai")
				sum += count(y)
			}
		}

	}
	a = append(a, sum)
	fmt.Println(sum)
	return a
}
func count(a [2]int) int {
	b := 0
	for _, v := range a {
		b += v
	}
	fmt.Println("ini b: ", b)
	return b
}

func main() {
	a := [][][2]int{
		{{5, 2}, {3, 2}}, {{10, 2}},
	}

	fmt.Println(CountProfit(a))
}

package main

import "fmt"

func CountProfit(data [][][2]int) []int {
	a := []int{}
	b := []int{}
	c := []int{}
	count := 0
	sum := 0
	if len(data) == 0 {

	} else {
		for i := 0; i < len(data); i++ {
			fmt.Println("1", i)
			hum := 0
			for j := 0; j < len(data[i]); j++ {
				fmt.Println("2", j)
				sum := 0
				for k := 0; k < len(data[i][j]); k++ {
					fmt.Println("3", k)
					if sum == 0 {
						sum += data[i][j][k]
					} else {
						hehe := sum - data[i][j][k]
						count = hehe
						fmt.Println("count: ", count)
						a = append(a, count)
					}
				}
				hum += count
				fmt.Println("hum: ", hum)
			}
			b = append(b, hum)
			sum += hum
			fmt.Println("sum: ", b)
		}
		c = append(c, sum)
		fmt.Println("sum: ", c)
	}

	return a // TODO: replace this
}

func main() {
	a := [][][2]int{
		{
			{1000, 500},
			{500, 150}}, {
			{600, 100},
			{800, 750},
		},
	}

	fmt.Println(CountProfit(a))
}

package main

import "fmt"

func CountingNumber(n int) float64 {
	sum := 1.0
	total := 0.0
	for sum <= float64(n) {
		fmt.Println("sum 1: ", sum)
		total += sum
		fmt.Println("total: ", total)
		sum += 0.5
		fmt.Println("sum 2: ", sum)
	}

	return total // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingNumber(5))
}

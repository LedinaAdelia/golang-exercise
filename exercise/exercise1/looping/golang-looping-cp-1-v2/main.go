package main

import "fmt"

func CountingNumber(n int) float64 {
	sum := 1.0
	total := 0.0
	for sum <= float64(n) {
		total += sum
		sum += 0.5
	}

	return total
}

func main() {
	fmt.Println(CountingNumber(5))
}
